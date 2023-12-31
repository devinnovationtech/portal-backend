package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type eventUcase struct {
	eventRepo      domain.EventRepository
	categories     domain.CategoryRepository
	tagsRepo       domain.TagRepository
	dataTagRepo    domain.DataTagRepository
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewEventUsecase will create new an eventUsecase object representation of domain.eventUsecase interface
func NewEventUsecase(repo domain.EventRepository, ctg domain.CategoryRepository, tags domain.TagRepository, dtags domain.DataTagRepository, user domain.UserRepository, timeout time.Duration) domain.EventUsecase {
	return &eventUcase{
		eventRepo:      repo,
		categories:     ctg,
		tagsRepo:       tags,
		dataTagRepo:    dtags,
		userRepo:       user,
		contextTimeout: timeout,
	}
}

// Private function to fill value of data tags
func (u *eventUcase) fillDataTags(c context.Context, data []domain.Event) ([]domain.Event, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the tags from the tags domain
	mapTags := map[int64][]domain.DataTag{}

	for _, eventTag := range data {
		mapTags[eventTag.ID] = []domain.DataTag{}
	}

	// Using goroutine to fetch the list tags
	chanTags := make(chan []domain.DataTag)
	for idx := range mapTags {
		eventID := idx
		g.Go(func() (err error) {
			res, err := u.dataTagRepo.FetchDataTags(ctx, eventID, domain.ConstEvent)
			chanTags <- res
			return
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanTags)
	}()

	for listTags := range chanTags {
		eventTags := []domain.DataTag{}
		copier.Copy(&eventTags, &listTags)
		if len(listTags) < 1 {
			continue
		}
		mapTags[listTags[0].DataID] = eventTags
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	for index, element := range data {
		if tags, ok := mapTags[element.ID]; ok {
			data[index].Tags = tags
		}
	}

	return data, nil
}

// Private function to fill value of detail data tags
func (u *eventUcase) fillDetailDataTags(c context.Context, data domain.Event) (domain.Event, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the tags from the tags domain
	mapTags := map[int64][]domain.DataTag{}

	mapTags[data.ID] = []domain.DataTag{}

	// Using goroutine to fetch the list tags
	chanTags := make(chan []domain.DataTag)
	for idx := range mapTags {
		eventID := idx
		g.Go(func() (err error) {
			res, err := u.dataTagRepo.FetchDataTags(ctx, eventID, domain.ConstEvent)
			chanTags <- res
			return
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanTags)
	}()

	for listTags := range chanTags {
		eventTags := []domain.DataTag{}
		copier.Copy(&eventTags, &listTags)
		if len(listTags) < 1 {
			continue
		}
		mapTags[listTags[0].DataID] = eventTags
	}

	if err := g.Wait(); err != nil {
		return data, err
	}

	if tags, ok := mapTags[data.ID]; ok {
		data.Tags = tags
	}

	return data, nil
}

// Private function to store tags
func (u *eventUcase) storeTags(ctx context.Context, eventID int64, tags []string, tx *sql.Tx) (err error) {
	for _, tagName := range tags {
		// 20 is max length of tags name
		tagName = helpers.Substr(tagName, 20)

		tag := &domain.Tag{
			Name: tagName,
		}

		// check tags already exists
		var checkTag domain.Tag
		checkTag, _ = u.tagsRepo.GetTagByName(ctx, tagName)
		tag.ID = checkTag.ID
		if checkTag.Name == "" {
			err = u.tagsRepo.StoreTag(ctx, tag, tx)
			if err != nil {
				return
			}
		}

		dataTag := &domain.DataTag{
			DataID:  eventID,
			TagID:   tag.ID,
			TagName: tagName,
			Type:    domain.ConstEvent,
		}
		err = u.dataTagRepo.StoreDataTag(ctx, dataTag, tx)
		if err != nil {
			return
		}
	}
	return
}

func (u *eventUcase) get(c context.Context, params *domain.Request) (res []domain.Event, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.eventRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = u.fillDataTags(ctx, res)

	if err != nil {
		return nil, 0, err
	}

	return
}

func filterByRoleAcces(au *domain.JwtCustomClaims, params *domain.Request) *domain.Request {

	if params.Filters == nil {
		params.Filters = map[string]interface{}{}
	}

	if !helpers.IsSuperAdmin(au) {
		params.Filters["unit_id"] = au.Unit.ID
	}

	return params
}

// Fetch will fetch data events
func (u *eventUcase) Fetch(c context.Context, au *domain.JwtCustomClaims, params *domain.Request) (res []domain.Event, total int64, err error) {

	return u.get(c, filterByRoleAcces(au, params))
}

// GetByID will find an object by given id
func (u *eventUcase) GetByID(c context.Context, id int64) (res domain.Event, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.eventRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	userData, err := u.userRepo.GetByID(ctx, res.CreatedBy.ID)
	if err != nil {
		return
	}
	res.CreatedBy = userData

	res, err = u.fillDetailDataTags(ctx, res)
	if err != nil {
		return
	}

	return
}

// GetByTitle will find an object by given name or title
func (u *eventUcase) GetByTitle(c context.Context, title string) (res domain.Event, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.eventRepo.GetByTitle(ctx, title)
	if err != nil {
		return
	}

	return
}

// Store will create you a new object, and store into database
func (u *eventUcase) Store(c context.Context, body *domain.StoreRequestEvent) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	tx, err := u.eventRepo.GetTx(c)
	if err != nil {
		return
	}

	body.CreatedAt = time.Now()
	body.UpdatedAt = time.Now()
	err = u.eventRepo.Store(ctx, body, tx)
	if err != nil {
		return
	}

	if err = u.storeTags(ctx, body.ID, body.Tags, tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

// Update will set up an update existing object
func (u *eventUcase) Update(c context.Context, id int64, body *domain.StoreRequestEvent) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	tx, err := u.eventRepo.GetTx(c)
	if err != nil {
		return
	}

	body.UpdatedAt = time.Now()

	err = u.eventRepo.Update(ctx, id, body, tx)

	err = u.dataTagRepo.DeleteDataTag(ctx, id, domain.ConstEvent, tx)
	if err != nil {
		return
	}

	if err = u.storeTags(ctx, id, body.Tags, tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

// Delete an object and destroy it from database
func (u *eventUcase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.eventRepo.Delete(ctx, id)
}

// AgendaPortal will get all object for portal endpoint
func (u *eventUcase) AgendaPortal(c context.Context, params *domain.Request) (res []domain.Event, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.eventRepo.AgendaPortal(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = u.fillDataTags(ctx, res)

	if err != nil {
		return nil, 0, err
	}

	return
}

// ListCalendar will get data event for calendar without paginate
func (u *eventUcase) ListCalendar(c context.Context, params *domain.Request) (res []domain.Event, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.eventRepo.ListCalendar(ctx, params)
	if err != nil {
		return nil, err
	}

	return
}

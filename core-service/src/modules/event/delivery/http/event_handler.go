package http

import (
	"net/http"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

// EventHandler is represented by domain.EventUsecase
type EventHandler struct {
	EventUcase domain.EventUsecase
}

// NewEventHandler will initialize the event endpoint
func NewEventHandler(e *echo.Group, r *echo.Group, us domain.EventUsecase) {
	handler := &EventHandler{
		EventUcase: us,
	}

	e.GET("/events", handler.Fetch)
	e.GET("/events/:id", handler.GetByID)
	e.GET("/events/calendar", handler.ListCalendar)
	e.POST("/events", handler.Store)
	e.DELETE("/events/:id", handler.Delete)
	e.PUT("/events/:id", handler.Update)
}

// Validate domain
func isRequestValid(m *domain.StoreRequestEvent) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Fetch will get events data
func (h *EventHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)

	listEvent, total, err := h.EventUcase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	listEventRes := []domain.ListEventResponse{}
	copier.Copy(&listEventRes, &listEvent)

	res := helpers.Paginate(c, listEventRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetByID will get event by given id
func (h *EventHandler) GetByID(c echo.Context) error {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	event, err := h.EventUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &event})
}

// ListCalendar ..
func (h *EventHandler) ListCalendar(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)

	listEvents, err := h.EventUcase.ListCalendar(ctx, &params)

	if err != nil {
		return nil
	}

	listEventCalendar := []domain.ListEventCalendarReponse{}
	copier.Copy(&listEventCalendar, &listEvents)

	return c.JSON(http.StatusOK, listEventCalendar)
}

// Store a new event ..
func (h *EventHandler) Store(c echo.Context) (err error) {
	var events domain.StoreRequestEvent
	var dt domain.DataTags
	err = c.Bind(&events)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&events); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = h.EventUcase.Store(ctx, &events, &dt)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, events)
}

// Delete an event ..
func (h *EventHandler) Delete(c echo.Context) (err error) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.EventUcase.Delete(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// Update an event ..
func (h *EventHandler) Update(c echo.Context) (err error) {
	var events domain.UpdateRequestEvent
	err = c.Bind(&events)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.EventUcase.Update(ctx, id, &events)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, events)
}

package http

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// PopUpBannerHandler ...
type PopUpBannerHandler struct {
	PUsecase domain.PopUpBannerUsecase
	apm      *utils.Apm
}

// NewPopUpBannerHandler will create a new PopUpBannerHandler
func NewPopUpBannerHandler(r *echo.Group, ucase domain.PopUpBannerUsecase, apm *utils.Apm) {
	handler := &PopUpBannerHandler{
		PUsecase: ucase,
		apm:      apm,
	}

	r.GET("/pop-up-banners", handler.Fetch)
	r.GET("/pop-up-banners/:id", handler.GetByID)
	r.POST("/pop-up-banners", handler.Store)
	r.DELETE("/pop-up-banners/:id", handler.Delete)
	r.PATCH("/pop-up-banners/:id/status", handler.UpdateStatus)
	r.PUT("/pop-up-banners/:id", handler.Update)
}

// Fetch will fetch the service-public
func (h *PopUpBannerHandler) Fetch(c echo.Context) error {
	// define requirements of request
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"status": c.QueryParam("status"),
	}

	auth := helpers.GetAuthenticatedUser(c)

	// usecase needed
	data, total, err := h.PUsecase.Fetch(ctx, auth, &params)
	if err != nil {
		return err
	}

	// re-presenting responses
	listPopUpResponse := []domain.ListPopUpBannerResponse{}
	for _, row := range data {
		// attach object response
		resp := domain.ListPopUpBannerResponse{
			ID:        row.ID,
			Title:     row.Title,
			Link:      row.Link,
			Status:    row.Status,
			IsLive:    row.IsLive,
			Duration:  row.Duration,
			StartDate: row.StartDate,
		}

		// un-marshalling object's string
		helpers.GetObjectFromString(row.Image.String, &resp.Image)

		// append element the end of slice
		listPopUpResponse = append(listPopUpResponse, resp)
	}

	res := helpers.Paginate(c, listPopUpResponse, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetByID will get pop up banner by given id
func (h *PopUpBannerHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	data, err := h.PUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// re-presenting responses
	res := domain.DetailPopUpBannerResponse{
		ID:          data.ID,
		Title:       data.Title,
		ButtonLabel: data.ButtonLabel,
		Link:        data.Link,
		Status:      data.Status,
		IsLive:      data.IsLive,
		Duration:    data.Duration,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
		UpdateAt:    data.UpdatedAt,
	}

	helpers.GetObjectFromString(data.Image.String, &res.Image)

	metaDesktop, _ := h.PUsecase.GetMetaDataImage(ctx, res.Image.Desktop) // for desktop
	metaMobile, _ := h.PUsecase.GetMetaDataImage(ctx, res.Image.Mobile)   // for mobile

	res.ImageMetaData.Desktop = metaDesktop
	res.ImageMetaData.Mobile = metaMobile

	metaRes := domain.MetaDetailPopUpBannerResponse{}
	copier.Copy(&metaRes, &res)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &metaRes})
}

// Store will store the pop up banner by given request body
func (h *PopUpBannerHandler) Store(c echo.Context) (err error) {
	req := new(domain.StorePopUpBannerRequest)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	ctx := c.Request().Context()
	err = h.PUsecase.Store(ctx, &auth, req)
	if err != nil {
		return err
	}

	res := domain.MessageResponse{
		Message: "successfully stored.",
	}

	return c.JSON(http.StatusCreated, res)
}

// Delete will delete the pop-up-banner by given id
func (h *PopUpBannerHandler) Delete(c echo.Context) (err error) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.PUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// UpdateStatus will update the pop-up-banners status by given request body
func (h *PopUpBannerHandler) UpdateStatus(c echo.Context) (err error) {
	body := new(domain.UpdateStatusPopUpBannerRequest)
	if err = c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err = validator.New().Struct(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()
	err = h.PUsecase.UpdateStatus(ctx, int64(reqID), body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully update status",
	})
}

// Update will update the pop-up-banners by given request body
func (h *PopUpBannerHandler) Update(c echo.Context) (err error) {
	body := new(domain.StorePopUpBannerRequest)
	if err = c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(body); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	au := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &au)

	ctx := c.Request().Context()
	err = h.PUsecase.Update(ctx, &au, int64(reqID), body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully updated.",
	})
}

func isRequestValid(ps interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}

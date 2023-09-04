package http

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

// PublicDocumentArchiveHandler is represented by domain.DocumentArchiveUsecase
type documentArchiveHandler struct {
	DocumentArchiveUcase domain.DocumentArchiveUsecase
}

// NewDocumentArchiveHandler will initialize the document archive endpoint
func NewDocumentArchiveHandler(r *echo.Group, us domain.DocumentArchiveUsecase) {
	handler := &documentArchiveHandler{
		DocumentArchiveUcase: us,
	}
	r.GET("/document-archives", handler.Fetch)
	r.POST("/document-archives", handler.Store)
	r.PUT("/document-archives/:id", handler.Update)
	r.DELETE("/document-archives/:id", handler.Delete)
	r.GET("/document-archives/:id", handler.GetByID)
	r.GET("/document-archives/tabs", handler.TabStatus)
	r.PATCH("/document-archives/:id/status", handler.UpdateStatus)
}

// Fetch for fetching document archive data
func (h *documentArchiveHandler) Fetch(c echo.Context) error {
	// init request by context
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Keyword = helpers.RegexCustomReplaceString(c, c.QueryParam("q"), "")
	params.Filters = map[string]interface{}{
		"category":   helpers.RegexReplaceString(c, c.QueryParam("cat"), ""),
		"categories": c.Request().URL.Query()["cat[]"],
		"status":     c.QueryParam("status"),
	}

	// getting data from usecase
	listDoc, total, err := h.DocumentArchiveUcase.Fetch(ctx, &params)
	if err != nil {
		return err
	}

	// preparing response
	listDocRes := []domain.ListDocumentArchive{}
	copier.Copy(&listDocRes, &listDoc)

	res := helpers.Paginate(c, listDocRes, total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *documentArchiveHandler) Store(c echo.Context) (err error) {
	req := new(domain.DocumentArchiveRequest)
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
	err = h.DocumentArchiveUcase.Store(ctx, req, auth.ID.String())
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	res := domain.MessageResponse{
		Message: "successfully stored.",
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *documentArchiveHandler) Delete(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseError{Message: domain.ErrNotFound.Error()})
	}

	ID := int64(id)
	ctx := c.Request().Context()

	if err = h.DocumentArchiveUcase.Delete(ctx, ID); err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.MessageResponse{
		Message: "Successfully deleted.",
	})
}

func (h *documentArchiveHandler) GetByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseError{Message: domain.ErrNotFound.Error()})
	}

	ID := int64(id)
	ctx := c.Request().Context()

	listDoc, err := h.DocumentArchiveUcase.GetByID(ctx, ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	listDocRes := domain.ListDocumentArchive{}
	copier.Copy(&listDocRes, &listDoc)

	return c.JSON(http.StatusOK, domain.ResultData{Data: listDocRes})
}

func (h *documentArchiveHandler) TabStatus(c echo.Context) error {
	ctx := c.Request().Context()

	// getting data from usecase
	res, err := h.DocumentArchiveUcase.TabStatus(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, domain.ResultData{
		Data: res,
	})
}

func (h *documentArchiveHandler) Update(c echo.Context) (err error) {
	req := new(domain.DocumentArchiveRequest)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))
	ID := int64(id)

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	ctx := c.Request().Context()
	err = h.DocumentArchiveUcase.Update(ctx, req, auth.ID.String(), ID)
	if err != nil {
		return err
	}

	res := domain.MessageResponse{
		Message: "successfully updated.",
	}

	return c.JSON(http.StatusOK, res)
}

func (h *documentArchiveHandler) UpdateStatus(c echo.Context) (err error) {
	req := new(domain.UpdateStatusDocumentArchiveRequest)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))
	ID := int64(id)

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	ctx := c.Request().Context()
	err = h.DocumentArchiveUcase.UpdateStatus(ctx, req, auth.ID.String(), ID)
	if err != nil {
		return err
	}

	res := domain.MessageResponse{
		Message: "successfully updated.",
	}

	return c.JSON(http.StatusOK, res)
}

func isRequestValid(ps interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}

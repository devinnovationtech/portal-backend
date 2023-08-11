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
	r.DELETE("/document-archives/:id", handler.Delete)
}

// Fetch for fetching document archive data
func (h *documentArchiveHandler) Fetch(c echo.Context) error {
	// init request by context
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"category": helpers.RegexReplaceString(c, c.QueryParam("cat"), ""),
		"status":   c.QueryParam("status"),
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
		return err
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

func isRequestValid(ps interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}

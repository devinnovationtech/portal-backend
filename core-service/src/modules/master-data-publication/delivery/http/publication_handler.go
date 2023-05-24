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
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/policies"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// MasterDataPublicationHandler ...
type MasterDataPublicationHandler struct {
	MdpUcase domain.MasterDataPublicationUsecase
	apm      *utils.Apm
}

// NewMasterDataPublicationHandler will create a new MasterDataPublicationHandler
func NewMasterDataPublicationHandler(r *echo.Group, sp domain.MasterDataPublicationUsecase, apm *utils.Apm) {
	handler := &MasterDataPublicationHandler{
		MdpUcase: sp,
		apm:      apm,
	}
	r.POST("/master-data-publications", handler.Store)
	r.GET("/master-data-publications", handler.Fetch)
	r.DELETE("/master-data-publications/:id", handler.Delete)
	r.GET("/master-data-publications/:id", handler.GetByID)
	r.GET("/master-data-publications/tabs", handler.TabStatus)
	r.PUT("/master-data-publications/:id", handler.Update)
}

func (h *MasterDataPublicationHandler) Store(c echo.Context) (err error) {
	// get a req context
	ctx := c.Request().Context()

	// bind a request body
	body, err := h.bindRequest(c)
	if err != nil {
		return
	}

	au := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &au)

	body.CreatedBy.ID = au.ID

	err = h.MdpUcase.Store(ctx, body)
	if err != nil {
		return err
	}

	res := map[string]interface{}{
		"message": "CREATED.",
	}

	return c.JSON(http.StatusCreated, res)
}

func isRequestValid(st interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *MasterDataPublicationHandler) bindRequest(c echo.Context) (body *domain.StoreMasterDataPublication, err error) {
	body = new(domain.StoreMasterDataPublication)
	if err = c.Bind(body); err != nil {
		return &domain.StoreMasterDataPublication{}, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(body); !ok {
		return &domain.StoreMasterDataPublication{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return
}

func (h *MasterDataPublicationHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"status": c.QueryParam("status"),
	}

	data, total, err := h.MdpUcase.Fetch(ctx, au, &params)
	if err != nil {
		return err
	}

	// represent responses to the client
	pubRes := []domain.ListMasterDataResponse{}
	for _, row := range data {
		res := domain.ListMasterDataResponse{
			ID:          row.ID,
			ServiceName: row.DefaultInformation.ServiceName,
			OpdName:     row.DefaultInformation.OpdName,
			ServiceUser: row.DefaultInformation.ServiceUser,
			Technical:   row.DefaultInformation.Technical,
			UpdatedAt:   row.UpdatedAt,
			Status:      row.Status,
		}

		pubRes = append(pubRes, res)
	}

	res := helpers.Paginate(c, pubRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// Delete will delete the master-data-publications by given id
func (h *MasterDataPublicationHandler) Delete(c echo.Context) (err error) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.MdpUcase.Delete(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// GetByID will get master data publication by given id
func (h *MasterDataPublicationHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	au := helpers.GetAuthenticatedUser(c)

	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	res, err := h.MdpUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	if !policies.AllowPublicationAccess(au, res) {
		return c.JSON(http.StatusForbidden, helpers.ResponseError{Message: domain.ErrForbidden.Error()})
	}

	// represent response to the client
	detailRes := domain.DetailPublicationResponse{}

	copier.Copy(&detailRes, &res)
	// un-marshalling json from string to object
	helpers.GetObjectFromString(res.DefaultInformation.Benefits.String, &detailRes.DefaultInformation.Benefits)
	helpers.GetObjectFromString(res.DefaultInformation.Facilities.String, &detailRes.DefaultInformation.Facilities)
	helpers.GetObjectFromString(res.ServiceDescription.Cover.String, &detailRes.ServiceDescription.Cover)
	helpers.GetObjectFromString(res.ServiceDescription.Images.String, &detailRes.ServiceDescription.Images)
	helpers.GetObjectFromString(res.ServiceDescription.TermsAndConditions.String, &detailRes.ServiceDescription.TermsAndConditions)
	helpers.GetObjectFromString(res.ServiceDescription.ServiceProcedures.String, &detailRes.ServiceDescription.ServiceProcedures)
	helpers.GetObjectFromString(res.ServiceDescription.OperationalTimes.String, &detailRes.ServiceDescription.OperationalTimes)
	helpers.GetObjectFromString(res.ServiceDescription.InfoGraphics.String, &detailRes.ServiceDescription.InfoGraphics)
	helpers.GetObjectFromString(res.ServiceDescription.Locations.String, &detailRes.ServiceDescription.Locations)
	helpers.GetObjectFromString(res.ServiceDescription.Application.Features.String, &detailRes.ServiceDescription.Application.Features)
	helpers.GetObjectFromString(res.ServiceDescription.Links.String, &detailRes.ServiceDescription.Links)
	helpers.GetObjectFromString(res.ServiceDescription.SocialMedia.String, &detailRes.ServiceDescription.SocialMedia)
	helpers.GetObjectFromString(res.AdditionalInformation.Keywords.String, &detailRes.AdditionalInformation.Keywords)
	helpers.GetObjectFromString(res.AdditionalInformation.FAQ.String, &detailRes.AdditionalInformation.FAQ)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &detailRes})
}

func (h *MasterDataPublicationHandler) TabStatus(c echo.Context) (err error) {
	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)
	params := helpers.GetRequestParams(c)

	tabs, err := h.MdpUcase.TabStatus(ctx, au, &params)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &tabs})
}

func (h *MasterDataPublicationHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ID := int64(reqID)
	body, err := h.bindRequest(c)
	if err != nil {
		return
	}

	err = h.MdpUcase.Update(ctx, body, ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	result := map[string]interface{}{
		"message": "UPDATED",
	}

	return c.JSON(http.StatusOK, result)
}

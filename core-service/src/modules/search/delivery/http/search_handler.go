package http

import (
	"net/http"
	"strings"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/labstack/echo/v4"
)

// SearchHandler ...
type SearchHandler struct {
	SUsecase domain.SearchUsecase
	Logger   *utils.Logrus
}

// NewSearchHandler will initialize the search/ resources endpoint
func NewSearchHandler(e *echo.Group, r *echo.Group, us domain.SearchUsecase, logger *utils.Logrus) {
	handler := &SearchHandler{
		SUsecase: us,
		Logger:   logger,
	}
	e.GET("/search", handler.FetchSearch)
	e.GET("/search/suggest", handler.SearchSuggestion)
}

// FetchSearch will fetch the content based on given params
func (h *SearchHandler) FetchSearch(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"domain":    c.Request().URL.Query()["domain[]"],
		"fuzziness": c.QueryParam("fuzziness"),
	}
	log := helpers.MapLog(c)
	log.Module = domain.SearchModule
	log.AdditionalInfo["searched_keywords"] = strings.ToLower(params.Keyword)

	listSearch, tot, aggs, err := h.SUsecase.Fetch(ctx, &params)
	if err != nil {
		h.Logger.Error(log, err)
		return err
	}

	res := helpers.Paginate(c, listSearch, tot, params)
	meta := res.Meta.(*domain.MetaData)
	meta.Aggregations = helpers.ESAggregate(aggs)

	disAllowedUserAgent := []string{
		"axios",
		"postman",
	}

	if !helpers.IsDisallowed(log.AdditionalInfo["user_agent"].(string), disAllowedUserAgent) { // TEMP: condition for mitigation non-organic search
		h.Logger.Info(log, "OK")
	}

	return c.JSON(http.StatusOK, res)
}

// SearchSuggestion ...
func (h *SearchHandler) SearchSuggestion(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"suggestions": c.QueryParam("q"),
	}

	if params.Filters["suggestions"] == "" {
		return c.JSON(http.StatusOK, "")
	}

	listSuggest, err := h.SUsecase.SearchSuggestion(ctx, &params)

	if err != nil {
		return err
	}

	if len(listSuggest) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(http.StatusOK, listSuggest)
}

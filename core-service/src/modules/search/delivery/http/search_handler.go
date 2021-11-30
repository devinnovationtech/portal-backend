package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
)

// SearchHandler ...
type SearchHandler struct {
	SUsecase domain.SearchUsecase
}

// NewSearchHandler will initialize the search/ resources endpoint
func NewSearchHandler(e *echo.Group, r *echo.Group, us domain.SearchUsecase) {
	handler := &SearchHandler{
		SUsecase: us,
	}
	e.GET("/search", handler.FetchSearch)
	e.GET("/search/suggestion", handler.SearchSuggestion)
}

// FetchSearch will fetch the content based on given params
func (h *SearchHandler) FetchSearch(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"domain": c.Request().URL.Query()["domain[]"],
	}

	listSearch, tot, aggs, err := h.SUsecase.Fetch(ctx, &params)
	if err != nil {
		return err
	}
	res := helpers.Paginate(c, listSearch, tot, params)
	meta := res.Meta.(*domain.MetaData)
	meta.Aggregations = helpers.ESAggregate(aggs)
	return c.JSON(http.StatusOK, res)
}

// SearchSuggestion ...
func (h *SearchHandler) SearchSuggestion(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"suggestion": c.QueryParam("q"),
	}

	if params.Filters["suggestion"] == "" {
		return c.JSON(http.StatusOK, "")
	}

	listSuggest, err := h.SUsecase.SearchSuggestion(ctx, &params)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listSuggest)
}

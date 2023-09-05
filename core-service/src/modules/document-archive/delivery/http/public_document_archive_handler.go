package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

// PublicDocumentArchiveHandler is represented by domain.DocumentArchiveUsecase
type PublicDocumentArchiveHandler struct {
	DocumentArchiveUcase domain.DocumentArchiveUsecase
	Logger               *utils.Logrus
}

// NewDocumentArchiveHandler will initialize the document archive endpoint
func NewPublicDocumentArchiveHandler(p *echo.Group, us domain.DocumentArchiveUsecase, logger *utils.Logrus) {
	handler := &PublicDocumentArchiveHandler{
		DocumentArchiveUcase: us,
		Logger:               logger,
	}
	p.GET("/document-archives", handler.Fetch)
}

// Fetch for fetching document archive data
func (h *PublicDocumentArchiveHandler) Fetch(c echo.Context) error {
	// init request by context
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"category": helpers.RegexReplaceString(c, c.QueryParam("cat"), ""),
		"status":   domain.DocumentArchivePublished,
	}
	log := helpers.MapLog(c)
	log.Module = domain.DocumentArchiveModule

	// getting data from usecase
	variant := unleash.GetVariant(domain.PortalDocumentArchive)
	var (
		listDoc []domain.DocumentArchive
		total   int64
		err     error
	)
	startTime := time.Now()
	if variant.Name == "query-without-goroutine" {
		listDoc, total, err = h.DocumentArchiveUcase.FetchWithoutGoRoutine(ctx, &params)
		log.AdditionalInfo["queries"] = "query-without-goroutine"
	} else {
		listDoc, total, err = h.DocumentArchiveUcase.Fetch(ctx, &params)
		log.AdditionalInfo["queries"] = "query-with-goroutine"
	}

	if err != nil {
		return err
	}
	difference := time.Now().Sub(startTime)
	log.AdditionalInfo["resp_elapsed_ms"] = fmt.Sprintf("%dms", difference.Milliseconds())
	log.Duration = int64(difference)

	h.Logger.Info(log, "OK")

	// preparing response
	listDocRes := []domain.ListDocumentArchive{}
	copier.Copy(&listDocRes, &listDoc)

	res := helpers.Paginate(c, listDocRes, total, params)

	return c.JSON(http.StatusOK, res)
}

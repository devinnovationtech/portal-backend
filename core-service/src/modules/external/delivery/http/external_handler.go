package http

import (
	"net/http"
	"strings"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/labstack/echo/v4"
)

type ExternalHandler struct {
	Logger *utils.Logrus
}

// NewExternalHandler will initialize the event endpoint
func NewExternalHandler(p *echo.Group, logger *utils.Logrus) {
	handler := &ExternalHandler{
		Logger: logger,
	}
	p.GET("/link-checker", handler.CheckLinkHandler)
}

// Fetch will get events data
func (h *ExternalHandler) CheckLinkHandler(c echo.Context) error {
	link := c.QueryParam("link")
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	}

	log := helpers.MapLog(c)
	log.Module = domain.ExternalModule
	log.AdditionalInfo["visited_external_links"] = link

	resp, err := http.Get(link)
	if err != nil {
		h.Logger.Error(log, err)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"link":  link,
			"valid": false,
			"ssl":   false,
		})
	}
	defer resp.Body.Close()

	valid := true
	ssl := false

	if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
		for _, cert := range resp.TLS.PeerCertificates {
			if cert.IsCA {
				ssl = true
				break
			}
		}
	}

	h.Logger.Info(log, "OK")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"link":  link,
		"valid": valid,
		"ssl":   ssl,
	})
}

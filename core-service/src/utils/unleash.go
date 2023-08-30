package utils

import (
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

func InitUnleash(cfg *config.Config) {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName(cfg.App.Name),
		unleash.WithUrl(cfg.Unleash.Url),
		unleash.WithCustomHeaders(http.Header{"Authorization": {cfg.Unleash.Token}}),
	)
}

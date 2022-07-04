package utils

import (
	"log"
	"os"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrzerolog"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
)

type Apm struct {
	NewRelic *newrelic.Application
}

func NewApm(cfg *config.Config) *Apm {
	return &Apm{
		NewRelic: initNewRelic(cfg),
	}
}

func initNewRelic(cfg *config.Config) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.NewRelic.AppName),
		newrelic.ConfigFromEnvironment(),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigLicense(cfg.NewRelic.License),
		newrelic.ConfigDistributedTracerEnabled(cfg.NewRelic.Enabled),
	)

	// logger := zerolog.New(os.Stdout)

	// txn := app.StartTransaction("My Transaction")
	// ctx := newrelic.NewContext(context.Background(), txn)

	// nrHook := nrzerolog.NewRelicHook{
	// 	App:     app,
	// 	Context: ctx,
	// }

	// nrLogger := logger.Hook(nrHook)

	// nrLogger.Info().Msg("A Log Message")

	// txn.End()

	logger := zerolog.New(os.Stdout)

	nrHook := nrzerolog.NewRelicHook{
		App: app,
	}

	nrLogger := logger.Hook(nrHook)

	nrLogger.Info().Msg("A Log Message")

	if err != nil {
		log.Fatal(err)
	}

	return app
}

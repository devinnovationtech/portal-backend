package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/cmd/server"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/middleware"
)

func main() {
	cfg := config.NewConfig()
	apm := utils.NewApm(cfg)
	conn := utils.NewDBConn(cfg)
	defer func() {
		err := conn.Mysql.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(nrecho.Middleware(apm.NewRelic))

	e.HTTPErrorHandler = server.ErrorHandler
	middL := middl.InitMiddleware(cfg)
	e.Use(middleware.CORSWithConfig(cfg.Cors))
	e.Use(middL.SENTRY)
	e.Use(middleware.Logger())

	// api v1
	v1 := e.Group("/v1")
	publicPath := v1.Group("/public")

	// restricted group
	restrictedPath := v1.Group("")
	restrictedPath.Use(middL.JWT)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Sentry.DSN,
		TracesSampleRate: cfg.Sentry.TracesSampleRate,
		Environment:      cfg.Sentry.Environment,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	e.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))

	timeoutContext := time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second

	// init repo category repo
	mysqlRepos := server.NewRepository(conn)
	usecases := server.NewUcase(cfg, conn, mysqlRepos, timeoutContext)
	server.NewHandler(v1, publicPath, restrictedPath, usecases)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}

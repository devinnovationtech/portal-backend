package server

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	_eventHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/delivery/http"
	_featuredProgramHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/delivery/http"
	_feedbackHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/delivery/http"
	_informationHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/delivery/http"
	_newsHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/delivery/http"
	_unitHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/delivery/http"
)

// NewHandler will create a new handler for the given usecase
func NewHandler(e *echo.Group, r *echo.Group, u *Usecases) {
	_newsHttpDelivery.NewNewsHandler(e, r, u.NewsUcase)
	_informationHttpDelivery.NewInformationHandler(e, r, u.InformationUcase)
	_unitHttpDelivery.NewUnitHandler(e, r, u.UnitUcase)
	_eventHttpDelivery.NewEventHandler(e, r, u.EventUcase)
	_feedbackHttpDelivery.NewFeedbackHandler(e, r, u.FeedbackUcase)
	_featuredProgramHttpDelivery.NewFeaturedProgramHandler(e, r, u.FeaturedProgramUcase)
}

// ErrorHandler ...
func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	sentry.CaptureException(err)
	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
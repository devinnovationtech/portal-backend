package helpers

import (
	"fmt"
	"time"

	_goLog "github.com/jabardigitalservice/golog/logger"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

func MapValue(arrMap []interface{}, key string) (value float64) {
	for _, m := range arrMap {
		el := m.(map[string]interface{})
		if el["key"] == key {
			return el["doc_count"].(float64)
		}
	}
	return
}

// MapUserInfo ...
func MapUserInfo(u domain.User) *domain.UserInfo {
	userinfo := domain.UserInfo{}
	copier.Copy(&userinfo, &u)
	return &userinfo
}

// Mapping log value use golog package
func MapLog(c echo.Context) *_goLog.LoggerData {
	protocol := "http"
	if c.Request().TLS != nil {
		protocol = "https"
	}

	return &_goLog.LoggerData{
		AdditionalInfo: map[string]interface{}{
			"http_host":         c.Request().Host,
			"http_addr":         protocol + "://" + c.Request().Host + c.Request().RequestURI,
			"http_method":       c.Request().Method,
			"http_proto":        c.Request().Proto,
			"http_schema":       protocol,
			"http_uri":          c.Request().RequestURI,
			"remote_addr":       c.Request().RemoteAddr,
			"resp_status":       c.Response().Status,
			"user_agent":        c.Request().UserAgent(),
			"resp_bytes_length": c.Request().ContentLength,
			"resp_elapsed_ms":   "0ms",
			"ts":                time.Now(),
		},
		Module:   "main",
		Category: "router",
		Duration: 0,
		Method:   fmt.Sprintf("[%s] %s", c.Request().Method, c.Request().URL),
	}
}

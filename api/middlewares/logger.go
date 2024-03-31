package middlewares

import (
	"authentication-test/api/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CustomLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return util.EndsWith(c.Path(), []string{"/assets*", "/*"})
		},
		Format: "[${time_rfc3339}] ${remote_ip}: ${status} ${method} ${uri} ${error}\n",
	})
}

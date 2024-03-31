package main

import (
	"authentication-test/api/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return util.EndsWith(c.Path(), []string{"/assets*", "/*"})
		},
		Format: "[${time_rfc3339}] ${remote_ip}: ${status} ${method} ${uri} ${error}\n",
	}))

	// Controllers

	// Serve react app
	e.Static("/assets", "./ui-dist/assets")
	e.Static("/", "./ui-dist")
	e.File("/", "ui-dist/index.html")
	//e.GET("*", func(c echo.Context) error {
	//	return c.File("ui-dist/index.html")
	//})

	e.Logger.Fatal(e.Start(":8080"))
}

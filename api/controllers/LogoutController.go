package controllers

import (
	"authentication-test/api/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		auth.ExpireCookies(c)
		return c.String(http.StatusOK, "Logged out")
	}
}

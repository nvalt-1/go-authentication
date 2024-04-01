package controllers

import (
	"authentication-test/api/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = auth.AccessTokenCookieName
		cookie.MaxAge = -1
		cookie.Path = "/"
		cookie.Value = ""
		cookie.HttpOnly = true
		cookie.SameSite = http.SameSiteStrictMode

		c.SetCookie(cookie)
		return c.String(http.StatusOK, "Logged out")
	}
}

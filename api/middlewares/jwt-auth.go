package middlewares

import (
	"authentication-test/api/auth"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTAuth() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: auth.GetClaims,
		SigningKey:    []byte(auth.GetJWTSecret()),
		TokenLookup:   "cookie:access-token",
		ErrorHandler:  auth.JWTErrorChecker,
	})
}

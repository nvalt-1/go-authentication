package middlewares

import (
	"authentication-test/api/auth"
	"authentication-test/api/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func TokenRefresher(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If the user is not authenticated (no user token data in the context), don't do anything.
		if c.Get("username") == nil {
			return next(c)
		}
		// Gets user token from the context.
		u := c.Get("username").(*jwt.Token)

		claims := u.Claims.(*auth.Claims)

		// We ensure that a new token is not issued until enough time has elapsed.
		// In this case, a new token will only be issued if the old token is within
		// 15 mins of expiry.
		if time.Unix(claims.ExpiresAt.Time.Unix(), 0).Sub(time.Now()) < 15*time.Minute {
			// Gets the refresh token from the cookie.
			rc, err := c.Cookie(auth.RefreshTokenCookieName)
			if err == nil && rc != nil {
				// Parses token and checks if it valid.
				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(auth.GetRefreshJWTSecret()), nil
				})
				if err != nil {
					if errors.Is(err, jwt.ErrSignatureInvalid) {
						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					// If everything is good, update tokens.
					_ = auth.GenerateTokensAndSetCookies(&models.User{
						Username: claims.Username,
					}, c)
				}
			}
		}

		return next(c)
	}
}

package controllers

import (
	"authentication-test/api/auth"
	"authentication-test/api/db"
	"authentication-test/api/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Load our "test" user.
		storedUser := db.LoadTestUser()
		u := new(models.User)

		// Parse the submitted data and fill the User struct with the data from the SignIn form.
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Compare the stored hashed password, with the hashed version of the password that was received.
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
			// If the two passwords don't match, return a 401 status.
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		// If password is correct, generate tokens and set cookies.
		err := auth.GenerateTokensAndSetCookies(storedUser, c)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		return c.String(http.StatusOK, fmt.Sprintf("You're in, %s", u.Username))
	}
}

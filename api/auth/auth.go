package auth

import (
	"authentication-test/api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	accessTokenCookieName = "access-token"
)

func GetJWTSecret() string {
	secret := os.Getenv("AUTH_TEST_SECRET")
	if secret == "" {
		log.Fatal("AUTH_TEST_SECRET not found in environment")
	}
	return secret
}

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetClaims(c echo.Context) jwt.Claims {
	return new(Claims)
}

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
func GenerateTokensAndSetCookies(user *models.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)

	return nil
}

func generateAccessToken(user *models.User) (string, time.Time, error) {
	// Declare the expiration time of the token (1h).
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

// Pay attention to this function. It holds the main JWT token generation logic.
func generateToken(user *models.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds.
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// Here we are creating a new cookie, which will store the valid JWT token.
func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

// Purpose of this cookie is to store the username.
func setUserCookie(user *models.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = user.Username
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

// JWTErrorChecker will be executed when user tries to access a protected path.
func JWTErrorChecker(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusUnauthorized)
}

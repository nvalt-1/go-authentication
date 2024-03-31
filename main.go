package main

import (
	"authentication-test/api/controllers"
	"authentication-test/api/middlewares"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	e := echo.New()

	// Middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middlewares.CustomLogger())

	// every api call requires authorization
	apiGroup := e.Group("/api")
	apiGroup.Use(middlewares.JWTAuth())
	apiGroup.Use(middlewares.TokenRefresher)

	// Controllers
	e.POST("/login", controllers.Login())

	// Serve react app
	e.Static("/assets", "./ui-dist/assets")
	e.Static("/", "./ui-dist")
	e.File("/", "ui-dist/index.html")

	e.Logger.Fatal(e.Start(":8080"))
}

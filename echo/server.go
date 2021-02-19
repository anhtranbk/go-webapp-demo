package echo

import (
	"webapp-demo/echo/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/registration", handlers.UserRegistration)
	auth.POST("/login", handlers.UserLogin)

	e.Logger.Fatal(e.Start(":8080"))
}

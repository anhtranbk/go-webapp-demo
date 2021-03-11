package echo

import (
	"webapp-demo/echo/handlers"
	"webapp-demo/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// register routes
	InitAuthRoutes(app, &handlers.AuthHandler{Service: service.NewMockAuthService()})

	app.Logger.Fatal(app.Start(":8080"))
}

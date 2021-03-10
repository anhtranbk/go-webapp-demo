package echo

import (
	"webapp-demo/echo/handlers"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(app *echo.Echo, handler *handlers.AuthHandler) {
	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/registration", handler.UserRegistration)
	auth.POST("/login", handler.UserLogin)
}

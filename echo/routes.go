package echo

import (
	"webapp-demo/echo/handler"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(app *echo.Echo, authHandler *handler.AuthHandler) {
	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/registration", authHandler.UserRegistration)
	auth.POST("/login", authHandler.UserLogin)
}

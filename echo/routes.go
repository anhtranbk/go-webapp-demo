package echo

import (
	"webapp-demo/echo/handler"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(router *echo.Group, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	auth.POST("/registration", authHandler.UserRegistration)
	auth.POST("/login", authHandler.UserLogin)
}

package echo

import (
	"webapp-demo/echo/handler"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(router *echo.Group, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	auth.POST("/signup", authHandler.SignUp)
	auth.POST("/signin", authHandler.SignIn)
}

package fiber

import (
	"webapp-demo/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/registration", authHandler.UserRegistration)
	auth.Post("/login", authHandler.UserLogin)
}

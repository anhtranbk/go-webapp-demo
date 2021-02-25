package fiber

import (
	"webapp-demo/fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoutes(app *fiber.App, handler *handlers.AuthHandler) {
	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/registration", handler.UserRegistration)
	auth.Post("/login", handler.UserLogin)
}

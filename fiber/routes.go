package fiber

import (
	"webapp-demo/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoutes(router fiber.Router, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	auth.Post("/registration", authHandler.UserRegistration)
	auth.Post("/login", authHandler.UserLogin)
}

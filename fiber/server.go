package fiber

import (
	"log"
	"webapp-demo/fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/registration", handlers.UserRegistration)
	auth.Post("/login", handlers.UserLogin)

	log.Fatal(app.Listen(":8080"))
}

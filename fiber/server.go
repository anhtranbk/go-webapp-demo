package fiber

import (
	"log"
	"webapp-demo/fiber/handlers"
	"webapp-demo/services"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	// init routes
	InitAuthRoutes(app, &handlers.AuthHandler{Service: services.NewMockAuthService()})

	log.Fatal(app.Listen(":8080"))
}

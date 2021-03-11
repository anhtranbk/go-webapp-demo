package fiber

import (
	"log"
	"webapp-demo/fiber/handlers"
	"webapp-demo/service"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	// init routes
	InitAuthRoutes(app, &handlers.AuthHandler{Service: service.NewMockAuthService()})

	log.Fatal(app.Listen(":8080"))
}

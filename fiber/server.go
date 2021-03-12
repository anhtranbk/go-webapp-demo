package fiber

import (
	"log"
	"webapp-demo/fiber/handler"
	"webapp-demo/service"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	// init routes
	InitAuthRoutes(app, &handler.AuthHandler{Service: service.NewMockAuthService()})

	log.Fatal(app.Listen(":8080"))
}

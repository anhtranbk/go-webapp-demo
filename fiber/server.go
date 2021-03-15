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
	v1 := app.Group("/v1")
	InitAuthRoutes(v1, &handler.AuthHandler{Service: service.NewMockAuthService()})

	log.Fatal(app.Listen(":8080"))
}

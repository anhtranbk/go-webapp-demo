package fiber

import (
	"context"
	"log"
	"webapp-demo/app"
	"webapp-demo/config"
	"webapp-demo/core"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	f := fiber.New()

	appCtx := core.AppContext{
		Context: context.TODO(),
		Config:  &config.Config{},
		Repo:    app.NewMockRepositories(),
	}

	InitRoutes(f, &appCtx)

	log.Fatal(f.Listen(":8080"))
}

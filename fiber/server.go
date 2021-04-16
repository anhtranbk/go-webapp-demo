package fiber

import (
	"context"
	"log"
	"webapp-demo/app/repository"
	"webapp-demo/config"
	"webapp-demo/core"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	appCtx := core.AppContext{
		Context: context.TODO(),
		Config:  &config.Config{},
		Repo:    repository.NewMockRepositories(),
	}

	InitRoutes(app, &appCtx)

	log.Fatal(app.Listen(":8080"))
}

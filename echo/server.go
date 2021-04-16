package echo

import (
	"context"
	"webapp-demo/app/repository"
	"webapp-demo/config"
	"webapp-demo/core"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	appCtx := core.AppContext{
		Context: context.TODO(),
		Config:  &config.Config{},
		Repo:    repository.NewMockRepositories(),
	}

	InitRoutes(e, &appCtx)

	e.Logger.Fatal(e.Start(":8080"))

}

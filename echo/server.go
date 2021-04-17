package echo

import (
	"context"
	"webapp-demo/app"
	"webapp-demo/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	appCtx := app.AppContext{
		Context: context.TODO(),
		Config:  &config.Config{},
		Repo:    app.NewMockRepositories(),
	}

	InitRoutes(e, &appCtx)

	e.Logger.Fatal(e.Start(":8080"))

}

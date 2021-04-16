package gin

import (
	"context"
	"webapp-demo/app/repository"
	"webapp-demo/config"
	"webapp-demo/core"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	appCtx := core.AppContext{
		Context: context.TODO(),
		Config:  &config.Config{},
		Repo:    repository.NewMockRepositories(),
	}

	InitRoutes(app, &appCtx)

	app.Run(":8080")
}

package gin

import (
	"context"
	"webapp-demo/app"
	"webapp-demo/config"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	g := gin.Default()

	appCtx := app.AppContext{
		Context: context.TODO(),
		Config:  &config.Config{},
		Repo:    app.NewMockRepositories(),
	}

	InitRoutes(g, &appCtx)

	g.Run(":8080")
}

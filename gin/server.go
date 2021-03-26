package gin

import (
	"context"
	"webapp-demo/config"
	"webapp-demo/core"
	"webapp-demo/gin/handler"
	"webapp-demo/repository"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	appCtx := &core.AppContext{
		Context: context.TODO(),
		Config:  &config.Config{},
		Repo:    repository.NewRepositories(),
	}

	// register routes
	v1 := app.Group("/v1")
	// InitAuthRoutes(v1, &handler.AuthHandler{Service: service.NewMockAuthService()})
	InitAuthRoutes(v1, handler.NewAuthHandler(appCtx))

	app.Run(":8080")
}

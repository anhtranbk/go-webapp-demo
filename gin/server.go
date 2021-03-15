package gin

import (
	"webapp-demo/gin/handler"
	"webapp-demo/service"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	// register routes
	v1 := app.Group("/v1")
	InitAuthRoutes(v1, &handler.AuthHandler{Service: service.NewMockAuthService()})

	app.Run(":8080")
}

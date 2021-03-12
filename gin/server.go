package gin

import (
	"webapp-demo/gin/handler"
	"webapp-demo/service"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	// register routes
	InitAuthRoutes(app, &handler.AuthHandler{Service: service.NewMockAuthService()})

	app.Run(":8080")
}

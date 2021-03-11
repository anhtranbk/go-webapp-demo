package gin

import (
	"webapp-demo/gin/handlers"
	"webapp-demo/service"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	// register routes
	InitAuthRoutes(app, &handlers.AuthHandler{Service: service.NewMockAuthService()})

	app.Run(":8080")
}

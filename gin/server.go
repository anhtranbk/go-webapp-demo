package gin

import (
	"webapp-demo/gin/handlers"
	"webapp-demo/services"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	// register routes
	InitAuthRoutes(app, &handlers.AuthHandler{Service: services.NewMockAuthService()})

	app.Run(":8080")
}

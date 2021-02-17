package gin

import (
	"webapp-demo/gin/handlers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	v1 := router.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/registration", handlers.UserRegistration)
	auth.POST("/login", handlers.UserLogin)

	router.Run(":8080")
}

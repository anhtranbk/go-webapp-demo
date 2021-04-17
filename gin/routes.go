package gin

import (
	"webapp-demo/app"
	"webapp-demo/gin/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine, appCtx *app.AppContext) {
	v1 := engine.Group("/v1")

	initAccountRoutes(v1, appCtx)
}

func initAccountRoutes(router *gin.RouterGroup, appCtx *app.AppContext) {
	auth := router.Group("/auth")
	// handler := &handler.AuthHandler{Service: service.NewMockAuthService()}
	handler := handler.NewAccountHandler(appCtx)

	auth.POST("/signup", handler.SignUp)
	auth.POST("/signin", handler.SignIn)
}

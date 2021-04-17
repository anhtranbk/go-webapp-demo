package echo

import (
	"webapp-demo/app"
	"webapp-demo/echo/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, appCtx *app.AppContext) {
	v1 := e.Group("/v1")

	initAccountRoutes(v1, appCtx)
}

func initAccountRoutes(router *echo.Group, appCtx *app.AppContext) {
	auth := router.Group("/auth")
	handler := handler.NewAccountHandler(appCtx)

	auth.POST("/signup", handler.SignUp)
	auth.POST("/signin", handler.SignIn)
}

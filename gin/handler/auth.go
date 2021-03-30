package handler

import (
	"net/http"
	"webapp-demo/core"
	"webapp-demo/dtos"
	"webapp-demo/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AppContext *core.AppContext
	Service    service.AuthService
}

func NewAuthHandler(ctx *core.AppContext) *AuthHandler {
	return &AuthHandler{
		AppContext: ctx,
		Service:    service.NewAuthService(&ctx.Context, *ctx.Repo),
	}
}

func (h AuthHandler) SignUp(ctx *gin.Context) {
	var signUp dtos.SignUpDto
	if err := ctx.ShouldBindJSON(&signUp); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserSignUp(signUp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h AuthHandler) SignIn(ctx *gin.Context) {
	var signIn dtos.SignInDto
	if err := ctx.ShouldBindJSON(&signIn); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserSignIn(signIn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

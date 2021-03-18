package handler

import (
	"net/http"
	"webapp-demo/dtos"
	"webapp-demo/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h AuthHandler) SignUp(ctx *gin.Context) {
	var signUp dtos.SignUpDto
	if err := ctx.ShouldBindJSON(&signUp); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserRegistration(signUp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &resp)
}

func (h AuthHandler) SignIn(ctx *gin.Context) {
	var signIn dtos.SignInDto
	if err := ctx.ShouldBindJSON(&signIn); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserLogin(signIn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &resp)
}

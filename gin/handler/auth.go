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
	var userReg dtos.SignUpDto
	if err := ctx.ShouldBindJSON(&userReg); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserRegistration(userReg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &resp)
}

func (h AuthHandler) SignIn(ctx *gin.Context) {
	var userLogin dtos.SignInDto
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserLogin(userLogin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &resp)
}

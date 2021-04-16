package handler

import (
	"net/http"
	"webapp-demo/app/dto"
	"webapp-demo/app/service"
	"webapp-demo/core"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	AppContext *core.AppContext
	Service    service.AccountService
}

func NewAccountHandler(appCtx *core.AppContext) *AccountHandler {
	return &AccountHandler{
		AppContext: appCtx,
		Service:    service.NewAccountService(appCtx),
	}
}

func (h AccountHandler) SignUp(ctx *gin.Context) {
	var signUp dto.SignUpDto
	if err := ctx.ShouldBindJSON(&signUp); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserSignUp(signUp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h AccountHandler) SignIn(ctx *gin.Context) {
	var signIn dto.SignInDto
	if err := ctx.ShouldBindJSON(&signIn); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.E{Error: err.Error()})
		return
	}

	resp, err := h.Service.UserSignIn(signIn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.E{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

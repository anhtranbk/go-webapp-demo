package handler

import (
	"net/http"
	"webapp-demo/core"
	"webapp-demo/dtos"
	"webapp-demo/service"

	echo "github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AppContext *core.AppContext
	Service    service.AuthService
}

func NewAuthHandler(appCtx *core.AppContext) *AuthHandler {
	return &AuthHandler{
		AppContext: appCtx,
		Service:    service.NewAuthService(appCtx),
	}
}

func (h AuthHandler) SignUp(c echo.Context) error {
	var signUp dtos.SignUpDto
	if err := c.Bind(&signUp); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignUp(signUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}

func (h AuthHandler) SignIn(c echo.Context) error {
	var signIn dtos.SignInDto
	if err := c.Bind(&signIn); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignIn(signIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}

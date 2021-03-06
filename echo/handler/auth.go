package handler

import (
	"net/http"
	"webapp-demo/app"
	"webapp-demo/app/user/dto"
	"webapp-demo/app/user/service"

	echo "github.com/labstack/echo/v4"
)

type AccountHandler struct {
	AppContext *app.AppContext
	Service    service.AccountService
}

func NewAccountHandler(appCtx *app.AppContext) *AccountHandler {
	return &AccountHandler{
		AppContext: appCtx,
		Service:    service.NewAccountService(appCtx),
	}
}

func (h AccountHandler) SignUp(c echo.Context) error {
	var signUp dto.SignUpDto
	if err := c.Bind(&signUp); err != nil {
		c.JSON(http.StatusBadRequest, dto.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignUp(signUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}

func (h AccountHandler) SignIn(c echo.Context) error {
	var signIn dto.SignInDto
	if err := c.Bind(&signIn); err != nil {
		c.JSON(http.StatusBadRequest, dto.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserSignIn(signIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}

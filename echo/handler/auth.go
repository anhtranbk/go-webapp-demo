package handler

import (
	"net/http"
	"webapp-demo/dtos"
	"webapp-demo/service"

	echo "github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h AuthHandler) SignUp(c echo.Context) error {
	var signUp dtos.SignUpDto
	if err := c.Bind(&signUp); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return err
	}

	resp, err := h.Service.UserRegistration(signUp)
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

	resp, err := h.Service.UserLogin(signIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, &resp)
	return nil
}

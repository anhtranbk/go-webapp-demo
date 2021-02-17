package handlers

import (
	"net/http"
	"webapp-demo/dtos"
	"webapp-demo/services"

	"github.com/gin-gonic/gin"
)

func UserRegistration(c *gin.Context) {
	var userReg dtos.UserRegistrationDto
	if err := c.ShouldBindJSON(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := services.UserRegistration(userReg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resp)
}

func UserLogin(c *gin.Context) {
	var userLogin dtos.UserLoginDto
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, dtos.E{Error: err.Error()})
		return
	}

	resp, err := services.UserLogin(userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.E{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &resp)
}

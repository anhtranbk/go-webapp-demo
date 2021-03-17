package service

import (
	"webapp-demo/dtos"
)

type AuthService interface {
	UserRegistration(userReg dtos.SignUpDto) (dtos.UserDto, error)
	UserLogin(userLogin dtos.SignInDto) (dtos.AccessTokenDto, error)
}

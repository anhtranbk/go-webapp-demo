package service

import (
	"webapp-demo/dtos"
)

type AuthService interface {
	UserRegistration(signUp dtos.SignUpDto) (*dtos.UserDto, error)
	UserLogin(signIn dtos.SignInDto) (*dtos.AccessTokenDto, error)
}

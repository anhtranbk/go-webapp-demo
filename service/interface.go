package service

import "webapp-demo/dto"

type AuthService interface {
	UserSignUp(signUp dto.SignUpDto) (*dto.UserDto, error)
	UserSignIn(signIn dto.SignInDto) (*dto.AccessTokenDto, error)
}

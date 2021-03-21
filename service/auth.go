package service

import (
	"context"
	"time"
	"webapp-demo/dtos"
	"webapp-demo/entity"
	"webapp-demo/pkg/errorx"
	"webapp-demo/repository"
)

type DefaultAuthService struct {
	context *context.Context
	repo    repository.Repositories
}

func NewDefaultAuthService(ctx *context.Context, repo repository.Repositories) *DefaultAuthService {
	return &DefaultAuthService{
		context: ctx,
		repo:    repo,
	}
}

func (s *DefaultAuthService) UserRegistration(signUp dtos.SignUpDto) (*dtos.UserDto, error) {
	repo := s.repo.UserRepo
	exist, err := repo.IsExist(signUp.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errorx.EmailExisted
	}

	user, err := repo.Create(&entity.User{
		UserName: signUp.UserName,
		Email:    signUp.Email,
		Password: signUp.Password,
	})
	if err != nil {
		return nil, err
	}

	return &dtos.UserDto{
		UserId:    uint64(user.UserId),
		UserName:  user.UserName,
		Email:     user.Email,
		Dob:       user.Dob,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *DefaultAuthService) UserLogin(signIn dtos.SignInDto) (*dtos.AccessTokenDto, error) {
	authRepo := s.repo.UserRepo
	refreshTokenRepo := s.repo.RefreshTokenRepo

	user, err := authRepo.FindByEmail(signIn.Email)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Password != signIn.Password {
		return nil, errorx.InvalidCredentials
	}

	refreshToken, err := refreshTokenRepo.Create("7c271bec2288b77c", time.Now().Add(30*24*time.Hour))
	if err != nil {
		return nil, err
	}

	return &dtos.AccessTokenDto{
		AccessToken:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1ODUwNjMyMTYsIm5iZiI6MTU4NTA2MzIxNiwianRpIj",
		TokenType:    "Bearer",
		ExpiredAt:    time.Now().Add(time.Minute * 60),
		RefreshToken: refreshToken.Token,
	}, nil
}

package service

import (
	"context"
	"time"
	"webapp-demo/config"
	"webapp-demo/core"
	"webapp-demo/dto"
	"webapp-demo/entity"
	"webapp-demo/pkg/errorx"
	"webapp-demo/repository"
)

type DefaultAuthService struct {
	config  *config.Config
	context context.Context
	repo    repository.Repositories
}

func NewAuthService(appCtx *core.AppContext) *DefaultAuthService {
	return &DefaultAuthService{
		config:  appCtx.Config,
		context: appCtx.Context,
		repo:    *appCtx.Repo,
	}
}

func (s *DefaultAuthService) UserSignUp(signUp dto.SignUpDto) (*dto.UserDto, error) {
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

	return &dto..UserDto{
		UserId:    user.UserId,
		UserName:  user.UserName,
		Email:     user.Email,
		Dob:       user.Dob,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *DefaultAuthService) UserSignIn(signIn dto.SignInDto) (*dto..AccessTokenDto, error) {
	authRepo := s.repo.UserRepo
	refreshTokenRepo := s.repo.RefreshTokenRepo

	user, err := authRepo.FindByEmail(signIn.Email)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Password != signIn.Password {
		return nil, errorx.InvalidCredentials
	}

	expiredAt := time.Now().Add(15 * 24 * time.Hour)
	refreshToken, err := refreshTokenRepo.Create("7c271bec2288b77c", expiredAt, user.UserId)
	if err != nil {
		return nil, err
	}

	return &dto..AccessTokenDto{
		AccessToken:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1ODUwNjMyMTYsIm5iZiI6MTU4NTA2MzIxNiwianRpIj",
		TokenType:    "Bearer",
		ExpiredAt:    time.Now().Add(time.Minute * 60),
		RefreshToken: refreshToken.Token,
	}, nil
}

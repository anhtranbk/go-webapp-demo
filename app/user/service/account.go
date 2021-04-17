package service

import (
	"context"
	"time"
	"webapp-demo/app"
	"webapp-demo/app/user/dto"
	"webapp-demo/app/user/entity"
	"webapp-demo/config"
	"webapp-demo/pkg/errorx"
	"webapp-demo/pkg/security/accesstoken"
	"webapp-demo/pkg/security/password"
	"webapp-demo/pkg/types"
	stringutil "webapp-demo/pkg/util/string"
	"webapp-demo/pkg/validation"
)

type DefaultAccountService struct {
	config  *config.Config
	context context.Context
	repo    app.Repositories
}

func NewAccountService(appCtx *app.AppContext) *DefaultAccountService {
	return &DefaultAccountService{
		config:  appCtx.Config,
		context: appCtx.Context,
		repo:    *appCtx.Repo,
	}
}

func (s *DefaultAccountService) UserSignUp(signUp dto.SignUpDto) (*dto.UserDto, error) {
	// validate input
	pwLen := len(signUp.Password)
	if pwLen < 6 || pwLen > 32 {
		return nil, errorx.MalformedPassword
	}
	if !validation.CheckEmailValid(signUp.Email) {
		return nil, errorx.MalformedEmail
	}

	repo := s.repo.UserRepo
	exist, err := repo.IsExist(signUp.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errorx.EmailExisted
	}

	pwHelper := password.NewPasswordHelper(s.config)
	hashedPassword, err := pwHelper.HashPassword(signUp.Password)
	if err != nil {
		return nil, err
	}

	user, err := repo.Create(&entity.User{
		UserName: signUp.UserName,
		Email:    signUp.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &dto.UserDto{
		UserId:    user.UserId,
		UserName:  user.UserName,
		Email:     user.Email,
		Dob:       user.Dob,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *DefaultAccountService) UserSignIn(signIn dto.SignInDto) (*dto.AccessTokenDto, error) {
	userRepo := s.repo.UserRepo
	refreshTokenRepo := s.repo.RefreshTokenRepo

	user, err := userRepo.FindByEmail(signIn.Email)
	if err != nil {
		return nil, err
	}

	pwHelper := password.NewPasswordHelper(s.config)
	if user == nil || !pwHelper.Matches(signIn.Password, user.Password) {
		return nil, errorx.InvalidCredentials
	}

	expiredAt := time.Now().Add(15 * 24 * time.Hour)
	tokenStr := stringutil.RandStringSeq(16)
	refreshToken, err := refreshTokenRepo.Create(tokenStr, expiredAt, user.UserId)
	if err != nil {
		return nil, err
	}

	// generate an accesstoken expires in one hour
	accessToken, err := accesstoken.GenerateToken(user.UserId, time.Hour,
		s.config.SecretKey, types.GenericMap{"email": user.Email})

	if err != nil {
		return nil, err
	}

	return &dto.AccessTokenDto{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		ExpiredAt:    time.Now().Add(time.Minute * 60),
		RefreshToken: refreshToken.Token,
	}, nil
}

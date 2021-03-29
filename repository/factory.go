package repository

import "webapp-demo/repository/mock"

func NewMockRepositories() *Repositories {
	return &Repositories{
		UserRepo:         mock.NewMockUserRepository(),
		RefreshTokenRepo: mock.NewMockRefreshTokenRepository(),
	}
}

package repository

func NewRepositories() *Repositories {
	return &Repositories{
		UserRepo:         NewMockUserRepository(),
		RefreshTokenRepo: NewMockRefreshTokenRepository(),
	}
}

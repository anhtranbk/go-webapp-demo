package repository

func NewMockRepositories() *Repositories {
	return &Repositories{
		UserRepo:         NewMockUserRepository(),
		RefreshTokenRepo: NewMockRefreshTokenRepository(),
	}
}

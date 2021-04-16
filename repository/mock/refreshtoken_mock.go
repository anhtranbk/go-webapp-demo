package mock

import (
	"math/rand"
	"time"
	"webapp-demo/entity"
)

type RefreshTokenRepo struct {
	tokens map[entity.ID]*entity.RefreshToken
}

func NewRefreshTokenRepo() *RefreshTokenRepo {
	return &RefreshTokenRepo{
		tokens: map[entity.ID]*entity.RefreshToken{},
	}
}

func (r *RefreshTokenRepo) Create(token string, expiredAt time.Time,
	userId entity.ID) (*entity.RefreshToken, error) {

	tokenObj := &entity.RefreshToken{
		TokenId:   entity.ID(rand.Int31n(999999)),
		Token:     token,
		ExpiredAt: expiredAt,
		IsActive:  true,
	}
	r.tokens[userId] = tokenObj
	return tokenObj, nil
}

func (r *RefreshTokenRepo) Update(token string, expiredAt time.Time,
	userId entity.ID) (*entity.RefreshToken, error) {

	tokenObj := r.tokens[userId]
	tokenObj.Token = token
	tokenObj.ExpiredAt = expiredAt
	return tokenObj, nil
}

func (r *RefreshTokenRepo) Deactivate(userId entity.ID) error {
	r.tokens[userId].IsActive = false
	return nil
}

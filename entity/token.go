package entity

import "time"

type RefreshToken struct {
	TokenId   ID
	Token     string
	ExpiredAt time.Time
	IsActive  bool
	UserId    ID
}

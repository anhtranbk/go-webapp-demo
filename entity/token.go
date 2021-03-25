package entity

import "time"

type RefreshToken struct {
	TokenId   EntityId
	Token     string
	ExpiredAt time.Time
	IsActive  bool
	UserId    EntityId
}

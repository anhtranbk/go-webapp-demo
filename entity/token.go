package entity

import "time"

type RefreshToken struct {
	UserId    EntityId
	Token     string
	ExpiredAt time.Time
	IsActive  bool
}

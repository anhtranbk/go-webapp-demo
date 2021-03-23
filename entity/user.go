package entity

import "time"

type EntityId uint64
type Gender string

const (
	Female Gender = "female"
	Male          = "male"
)

type User struct {
	UserId    EntityId
	UserName  string
	Password  string
	Email     string
	FullName  string
	Dob       string
	Gender    Gender
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
	IsDeleted bool
}

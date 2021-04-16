package entity

import "time"

type ID int64
type Gender string

const (
	Female Gender = "female"
	Male          = "male"
)

type User struct {
	UserId    ID
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

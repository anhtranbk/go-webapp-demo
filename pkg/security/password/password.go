package password

import (
	"webapp-demo/config"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHelper interface {
	HashPassword(rawPassword string) (string, error)
	Matches(rawPassword string, hashedPassword string) bool
}

// factory method, allow configure password helper from configuration
func NewPasswordHelper(c *config.Config) PasswordHelper {
	return NewBcryptPasswordHelper()
}

type BcryptPasswordHelper struct {
}

// constructor
func NewBcryptPasswordHelper() *BcryptPasswordHelper {
	return &BcryptPasswordHelper{}
}

func (h *BcryptPasswordHelper) HashPassword(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	return string(bytes), err
}

func (h *BcryptPasswordHelper) Matches(rawPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}

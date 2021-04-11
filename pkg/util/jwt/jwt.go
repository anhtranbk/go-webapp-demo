package jwtutil

import (
	"fmt"
	"webapp-demo/pkg/types"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.(jwt.MapClaims))
	return token.SignedString([]byte(secret))
}

func ParseAndValidate(secret string, tokenString string) (types.GenericMap, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return types.GenericMap(claims), nil
	}
	return nil, err
}

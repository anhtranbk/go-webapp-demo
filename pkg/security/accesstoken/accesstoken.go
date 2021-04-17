package accesstoken

import (
	"strconv"
	"time"
	"webapp-demo/app/user/entity"
	"webapp-demo/pkg/types"
	jwtutil "webapp-demo/pkg/util/jwt"
	maputil "webapp-demo/pkg/util/map"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateToken(userId entity.ID, expireIn time.Duration,
	secret string, customClaims types.GenericMap) (string, error) {

	claims := jwt.MapClaims{
		"sub": strconv.FormatInt(int64(userId), 10),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(expireIn).Unix(),
		"jti": uuid.NewString(),
	}
	maputil.UpdateGenericMap(types.GenericMap(claims), customClaims)

	return jwtutil.GenerateToken(secret, claims)
}

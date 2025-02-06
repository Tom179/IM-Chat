package ctxdata

import "github.com/golang-jwt/jwt"

const Identify = "im-chat"

func GetJwtToken(secretKey string, iat, expireSeconds int64, uid string) (string, error) {
	//iat：int64签发时间
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expireSeconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

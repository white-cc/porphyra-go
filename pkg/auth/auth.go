package auth

import (
	"encoding/hex"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwt_key = []byte(os.Getenv("jwt_key"))

type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}

func CreateToken(name string, pass []byte) (string, error) {
	password := hex.EncodeToString(pass)
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: name,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwt_key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})
	if err != nil {
		return nil, err
	}
	//使用断言解析转换
	return tokenClaims.Claims.(*Claims), nil
}

package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "secret"
)

func CreateToken(Issuer string) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    Issuer,
		ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
	})
	token, err = claims.SignedString([]byte(SecretKey))
	return
}

func VerifyToken(tokenString string) (bool, string) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	claims := token.Claims.(*jwt.StandardClaims)
	return err == nil, claims.Issuer
}


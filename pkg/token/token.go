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


func VerifyToken(token string)(bool){
	_, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	return err == nil
}
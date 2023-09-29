package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"simpleCSS2/core/define"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string) (string, error) {
	//id

	//identity
	//name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString(define.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

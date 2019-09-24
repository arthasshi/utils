package utils

import (
	"github.com/dgrijalva/jwt-go"
)

func GenJwt(uid string) (string, error) {
	uidByte := []byte(uid)
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "sspass-cms",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(uidByte)
	return ss, err
}

package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("sdfghjkhgvfjbhgfdfdsdfghjwertyu")

func GenerateToken(publicKey string) (string, error) {
	claims := jwt.MapClaims{
		"publicKey": publicKey,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)
}

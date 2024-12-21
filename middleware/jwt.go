package middleware

import (
	"errors"
	"go/token"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("sdfghjkhgvfjbhgfdfdsdfghjwertyu")

var jwtError = errors.New("Invalid token.")

func GenerateToken(publicKey string) (string, error) {
	claims := jwt.MapClaims{
		"publicKey": publicKey,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwtError
		}
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

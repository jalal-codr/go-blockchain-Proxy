package services

import (
	models "proxy/Models"
	"proxy/middleware"
)

func SignIn(publicKey string) (string, error) {
	user, err := models.GetUser(publicKey)
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateToken(user.Publickey)
	if err != nil {
		return "", err
	}
	return token, nil
}

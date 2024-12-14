package services

import (
	"fmt"
	models "proxy/Models"
	"proxy/types"
	"proxy/utils"
)

func CreateUser(hash string) (string, error) {
	privateKey, err := utils.GenerateRandomBase64String(32)
	if err != nil {
		fmt.Println("Error creating private key", err)
		return "", err
	}

	publicKey, err := utils.GenerateRandomBase64String(32)
	if err != nil {
		fmt.Println("Error creating public key", err)
		return "", err
	}

	newUser := new(types.User)
	newUser.Publickey = publicKey

	encryptedHash, err := utils.EncryptString(hash, privateKey)
	if err != nil {
		fmt.Println("Error encrypting hash", err)
		return "", err
	}
	newUser.Hash = encryptedHash

	encryptedPrivateKey, err := utils.EncryptString(privateKey, publicKey)
	if err != nil {
		fmt.Println("Error encrypting privateKey", err)
		return "", err
	}
	newUser.PrivateKey = encryptedPrivateKey

	err = models.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error saving user", err)
		return "", err
	}

	return publicKey, nil
}

func GetUserHash(user types.User) (string, error) {
	privateKey, err := utils.DecryptString(user.PrivateKey, user.Publickey)
	if err != nil {
		return "", err
	}
	hash, err := utils.DecryptString(user.Hash, privateKey)
	if err != nil {
		return "", err
	}
	return hash, nil
}

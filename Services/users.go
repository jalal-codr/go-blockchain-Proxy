package services

import (
	"fmt"
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

	return publicKey, nil
}

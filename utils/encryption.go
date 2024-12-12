package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// Encrypt a string using AES encryption
func EncryptString(data string) (string, error) {
	key := os.Getenv("ENCRYPTION_KEY")
	// Convert the key to a 32-byte array
	keyBytes := []byte(key)
	if len(keyBytes) != 32 {
		return "", fmt.Errorf("encryption key must be 32 bytes")
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("error creating cipher block: %w", err)
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("error generating IV: %w", err)
	}

	// Encrypt the data using Cipher Block Chaining (CBC) mode
	ciphertext := make([]byte, len(data))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, []byte(data))

	// Prepend the IV to the ciphertext (needed for decryption)
	result := append(iv, ciphertext...)

	// Encode the encrypted data as a base64 string
	return base64.StdEncoding.EncodeToString(result), nil
}

func DecryptString(encrypted string) (string, error) {
	key := os.Getenv("ENCRYPTION_KEY")
	keyBytes := []byte(key)
	if len(keyBytes) != 32 {
		return "", fmt.Errorf("encryption key must be 32 bytes")
	}

	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("error decoding base64: %w", err)
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("error creating cipher block: %w", err)
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

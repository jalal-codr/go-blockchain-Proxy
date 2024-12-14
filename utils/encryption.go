package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func EncryptString(data, privateKey string) (string, error) {
	// Get the base64-encoded key from the environment variable
	keyBase64 := privateKey
	if keyBase64 == "" {
		return "", fmt.Errorf("ENCRYPTION_KEY not set")
	}

	// Decode the base64 key
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 key: %w", err)
	}

	// Ensure the key is exactly 32 bytes
	if len(key) != 32 {
		return "", fmt.Errorf("encryption key must be 32 bytes")
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
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

func DecryptString(encrypted, privateKey string) (string, error) {
	// Get the base64-encoded key from the environment variable
	keyBase64 := privateKey
	if keyBase64 == "" {
		return "", fmt.Errorf("ENCRYPTION_KEY not set")
	}

	// Decode the base64 key
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 key: %w", err)
	}

	// Ensure the key is exactly 32 bytes
	if len(key) != 32 {
		return "", fmt.Errorf("encryption key must be 32 bytes")
	}

	// Decode the base64-encoded encrypted data
	cipherData, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 encrypted data: %w", err)
	}

	// Extract the IV (first AES.BlockSize bytes) and ciphertext
	if len(cipherData) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := cipherData[:aes.BlockSize]
	ciphertext := cipherData[aes.BlockSize:]

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creating cipher block: %w", err)
	}

	// Decrypt the data using Cipher Block Chaining (CBC) mode
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	// Return the plaintext as a string
	return string(plaintext), nil
}

func GenerateRandomBase64String(length int) (string, error) {
	// Generate a random byte slice
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Encode the bytes into a Base64 string
	return base64.StdEncoding.EncodeToString(bytes), nil
}

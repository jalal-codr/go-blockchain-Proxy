package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateBlock(key string) (string, error) {
	url := "http://localhost:8080/createBlock"

	// Example JSON data to send
	data := map[string]interface{}{
		"data": key,
	}

	// Marshal the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON data: %w", err)
	}

	// Create a new POST request with the JSON payload
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the response JSON to extract blockHash
	var response map[string]string
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	blockHash, ok := response["blockHash"]
	if !ok {
		return "", fmt.Errorf("blockHash not found in response")
	}

	// Encrypt the blockHash
	publicKey, err := CreateUser(blockHash)
	if err != nil {
		return "", fmt.Errorf("error encrypting blockHash: %w", err)
	}

	// Return the encrypted blockHash
	return publicKey, nil
}

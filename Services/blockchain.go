package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	models "proxy/Models"

	"github.com/gorilla/websocket"
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

func GetBalance(publicKey string) (interface{}, error) {
	user, err := models.GetUser(publicKey)
	if err != nil {
		return 0, fmt.Errorf("Error fetching user: %w", err)
	}
	hash, err := GetUserHash(user)
	if err != nil {
		return 0, fmt.Errorf("error fetching hash: %w", err)
	}

	url := "http://localhost:8080/getBalance"

	// Example JSON data to send
	data := map[string]interface{}{
		"hash": hash,
	}

	// Marshal the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, fmt.Errorf("error marshaling JSON data: %w", err)
	}

	// Create a new POST request with the JSON payload
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the response JSON to extract blockHash
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	balance, ok := response["balance"]
	if !ok {
		return 0, fmt.Errorf("balance not found")
	}
	// Return the encrypted blockHash
	return balance, nil
}

func mining(hash string) {
	serverURL := "ws://localhost:8080/ws"

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connected to WebSocket server")

	for {
		var data struct {
			Hash string `json:"hash"`
		}
		data.Hash = hash
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error serializing to JSON:", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		_, response, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Print the response from the server
		fmt.Printf("Received response: %s\n", response)
	}
}

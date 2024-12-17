package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	models "proxy/Models"
	services "proxy/Services"
	"proxy/types"

	"github.com/gorilla/websocket"
)

func CreateBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invald request method", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request bodys", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data types.NewBlock
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Ivalid Json format", http.StatusBadRequest)
		return
	}
	newBlock, err := services.CreateBlock(data.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// res, err := utils.DecryptString(newBlock)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message":   "Block created",
		"publicKey": newBlock,
	}
	json.NewEncoder(w).Encode(response)
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invald request method", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request bodys", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data types.UserFetch
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Ivalid Json format", http.StatusBadRequest)
		return
	}
	balance, err := services.GetBalance(data.Publickey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Request Succesful",
		"balance": balance,
	}
	json.NewEncoder(w).Encode(response)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection", err)
	}
	defer conn.Close()
	for {
		_, publicKey, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message", err)
			break
		}
		var data struct {
			Publickey string `json:"publicKey"`
		}
		if err := json.Unmarshal(publicKey, &data); err != nil {
			fmt.Println("Invalid json format", err)
		}
		user, err := models.GetUser(data.Publickey)
		if err != nil {
			fmt.Println("Error fetching user", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid publicKey"))
			return
		}
		hash, err := services.GetUserHash(user)
		if err != nil {
			fmt.Println("Error fetching hash", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid publicKey"))
			return
		}
		msgChan := make(chan []byte)
		errChan := make(chan error)
		go services.Mining(hash, msgChan, errChan)
		for {
			select {
			case msg := <-msgChan:
				conn.WriteMessage(websocket.TextMessage, msg)
			case err := <-errChan:
				fmt.Println("Error:", err)
				return
			}
		}

	}
}

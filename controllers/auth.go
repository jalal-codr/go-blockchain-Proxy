package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	services "proxy/Services"
	"proxy/types"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
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
	result, err := services.SignIn(data.Publickey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := result
	json.NewEncoder(w).Encode(response)
}

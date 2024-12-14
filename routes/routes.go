package routes

import (
	"fmt"
	"log"
	"net/http"
	"proxy/controllers"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Define routes and their handlers
	mux.HandleFunc("/", controllers.HelloHandler)
	mux.HandleFunc("/createBlock", controllers.CreateBlock)

	// mux.HandleFunc("/ws/mining", controllers.WebsocketConnection)

	// mux.HandleFunc("/transferToken", controllers.TransferToken)

	mux.HandleFunc("/getBalance", controllers.GetBalance)

	return mux
}

func StartServer() {
	router := InitRoutes()

	// Start the server
	fmt.Println("Starting server on :7000")
	if err := http.ListenAndServe(":7000", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"proxy/routes"
)

func StartServer() {
	router := routes.InitRoutes()

	// Start the server
	fmt.Println("Starting server on :7000")
	if err := http.ListenAndServe(":7000", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

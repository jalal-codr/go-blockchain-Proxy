package main

import (
	"log"
	"os"
	"proxy/jobs"
	"proxy/routes"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, running without .env")
		}
	}

	go jobs.RunJobs()
	routes.StartServer()
}

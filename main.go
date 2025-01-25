package main

import (
	"fmt"
	"log"
	"os"
	models "proxy/Models"
	"proxy/routes"
	"proxy/templates"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, running without .env")
		}
	}

	go models.InitDb()
	err := templates.InitializeTemplates()
	if err != nil {
		fmt.Println(err)
	}
	// go jobs.RunJobs()
	routes.StartServer()
}

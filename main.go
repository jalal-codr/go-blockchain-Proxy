package main

import (
	"proxy/jobs"
	"proxy/routes"
)

func main() {
	go jobs.RunJobs()
	routes.StartServer()
}

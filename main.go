package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	api "github.com/darksuei/kubeRPC-sidecar/api"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3500" // Default to 3500 if PORT is not set
	}

	// Health API
	http.HandleFunc("/health", api.Health)

	log.Println("Sidecar running on port: ", port)

	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Println("Error starting sidecar:", err)
	}
}
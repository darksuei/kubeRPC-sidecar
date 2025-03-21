package main

import (
	"net/http"
	"github.com/joho/godotenv"

	api "github.com/darksuei/kubeRPC/api"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3500" // Default to 3500 if PORT is not set
	}

	// Health API
	http.HandleFunc("/health", api.Health)

	fmt.Println("Sidecar running on port: ", port)

	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		fmt.Println("Error starting sidecar:", err)
	}
}
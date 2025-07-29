package main

import (
	"log"
	"net/http"

	"caas-aks-api-go/api"
)

// @title CaaS AKS API
// @version 1.0
// @description API to manage Azure AKS clusters
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	r := api.SetupRouter()
	log.Println("🚀 AKS API server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

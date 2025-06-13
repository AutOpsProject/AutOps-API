package main

import (
	"log"
	"net/http"

	"github.com/AutOpsProject/AutOps-API/internal/api"
)

func main() {
	router := api.SetupRouter()
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}

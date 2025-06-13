package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()
	// r.HandleFunc("/project", handler.CreateProject).Methods("POST")
	return r
}

package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// InitializeRoutes : All the endpoints/routes
func InitializeRoutes() {
	addr := ":8080"
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/createUserAndWorkspace", createUserAndWorkspaceHandler).Methods(http.MethodPost)
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

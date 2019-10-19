package controllers

import (
	"fmt"
	"net/http"
)

// CreateWorkspace : Create user workspace
func (server *Server) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}

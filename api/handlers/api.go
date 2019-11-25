package handlers

import "net/http"

// Endpoints : All the endpoints
func Endpoints() {
	http.HandleFunc("/api/v1/createUserAndWorkspace", createUserAndWorkspaceHandler)
}

package handlers

import (
	"encoding/json"
	"net/http"
)

// Request : Signup request struct
type Request struct {
	Email     string `json:"email"`
	Workspace string `json:"workspace"`
}

func createUserAndWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

}

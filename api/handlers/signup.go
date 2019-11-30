package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZeeshanTamboli/slack-clone-services/database"
)

// Request : Signup request struct
type Request struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Workspace string `json:"workspace"`
}

func createUserAndWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	res, err := database.DBCon.Exec(`insert into users (first_name, last_name, email) values ($1, $2, $3)`, request.FirstName, request.LastName, request.Email)
	if err != nil {
		panic(err)
	}
	numOfRowsInserted, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of rows inserted : %v", numOfRowsInserted)
}

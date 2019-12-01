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

// Response : Signup response struct
type Response struct {
	Success bool `json:"success"`
}

func createUserAndWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	response := Response{}
	var numOfRowsInserted int64

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	var userID int
	err = database.DBCon.QueryRow(`select id from users where email=$1`, request.Email).Scan(&userID)
	if userID == 0 {
		err = database.DBCon.QueryRow(`insert into users (first_name, last_name, email) values ($1, $2, $3) returning id`, request.FirstName, request.LastName, request.Email).Scan(&userID)
		if err != nil {
			panic(err)
		}
	}

	var workspaceID int
	err = database.DBCon.QueryRow(`select id from workspaces where name=$1`, request.Workspace).Scan(&workspaceID)
	if workspaceID == 0 {
		err = database.DBCon.QueryRow(`insert into workspaces (name) values ($1) returning id`, request.Workspace).Scan(&workspaceID)
		if err != nil {
			panic(err)
		}
	}

	usersAndworkspacesRes, err := database.DBCon.Exec(`insert into users_workspaces (user_id, workspace_id) values ($1, $2)`, userID, workspaceID)
	if err != nil {
		panic(err)
	}
	numOfRowsInserted, err = usersAndworkspacesRes.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of rows inserted in 'users_workspaces' table : %v \n", numOfRowsInserted)

	response.Success = true
	responseJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

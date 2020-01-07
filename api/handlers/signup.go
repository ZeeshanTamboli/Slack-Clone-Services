package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ZeeshanTamboli/slack-clone-services/api/responses"

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
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Cannot decode request JSON"))
		return
	}

	var userID int
	err = database.DBCon.QueryRow(`select id from users where email=$1`, request.Email).Scan(&userID)
	if userID == 0 {
		err = database.DBCon.QueryRow(`insert into users (first_name, last_name, email) values ($1, $2, $3) returning id`, request.FirstName, request.LastName, request.Email).Scan(&userID)
		if err != nil {
			responses.ERROR(w, http.StatusConflict, errors.New("Cannot insert user data"))
			return
		}
	}

	var workspaceID int
	err = database.DBCon.QueryRow(`select id from workspaces where name=$1`, request.Workspace).Scan(&workspaceID)
	if workspaceID == 0 {
		err = database.DBCon.QueryRow(`insert into workspaces (name) values ($1) returning id`, request.Workspace).Scan(&workspaceID)
		if err != nil {
			responses.ERROR(w, http.StatusConflict, errors.New("Cannot insert workspace data"))
			return
		}
	}

	usersAndworkspacesRes, err := database.DBCon.Exec(`insert into users_workspaces (user_id, workspace_id) values ($1, $2)`, userID, workspaceID)
	if err != nil {
		responses.ERROR(w, http.StatusConflict, errors.New("Cannot insert duplicate data into users_workspaces"))
		return
	}
	numOfRowsInserted, err = usersAndworkspacesRes.RowsAffected()
	if err != nil {
		responses.ERROR(w, http.StatusConflict, errors.New("Rows affected throwed an error"))
		return
	}
	fmt.Printf("Number of rows inserted in 'users_workspaces' table : %v \n", numOfRowsInserted)

	response.Success = true
	responses.JSON(w, http.StatusOK, response)
}

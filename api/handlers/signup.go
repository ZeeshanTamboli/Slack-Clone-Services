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
type signupRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Workspace string `json:"workspace"`
}

// Response : Signup response struct
type signupResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func createUserAndWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	request := signupRequest{}
	response := signupResponse{}
	var numOfRowsInserted int64

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Cannot decode request JSON"))
		return
	}

	// Check if there is a owner for the workspace
	var ownerID int
	database.DBCon.QueryRow(`select owner_id from workspaces where name=$1`, request.Workspace).Scan((&ownerID))
	if ownerID != 0 {
		responses.ERROR(w, http.StatusConflict, errors.New("Workspace already exists"))
		return
	}

	var userID int
	database.DBCon.QueryRow(`select id from users where email=$1`, request.Email).Scan(&userID)
	if userID == 0 {
		err = database.DBCon.QueryRow(`insert into users (first_name, last_name, email) values ($1, $2, $3) returning id`, request.FirstName, request.LastName, request.Email).Scan(&userID)
		if err != nil {
			responses.ERROR(w, http.StatusConflict, errors.New("Cannot insert user data"))
			return
		}
	}

	var workspaceID int
	database.DBCon.QueryRow(`select id from workspaces where name=$1`, request.Workspace).Scan(&workspaceID)
	if workspaceID == 0 {
		err = database.DBCon.QueryRow(`insert into workspaces (name, owner_id) values ($1, $2) returning id`, request.Workspace, userID).Scan(&workspaceID)
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
		responses.ERROR(w, http.StatusConflict, errors.New("RowsAffected throwed an error"))
		return
	}
	fmt.Printf("Number of rows inserted in 'users_workspaces' table : %v \n", numOfRowsInserted)

	response.Success = true
	responses.JSON(w, http.StatusOK, response)
}

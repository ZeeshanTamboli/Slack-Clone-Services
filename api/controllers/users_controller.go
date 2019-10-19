package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ZeeshanTamboli/slack-clone-services/api/responses"

	"github.com/ZeeshanTamboli/slack-clone-services/api/models"
)

// CreateWorkspace : Create user workspace
func (server *Server) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	var err error
	user := models.User{}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

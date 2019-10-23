package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ZeeshanTamboli/slack-clone-services/api/responses"

	"github.com/ZeeshanTamboli/slack-clone-services/api/models"
)

// CreateWorkspace : Create user workspace
func (server *Server) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	workspace := models.Workspace{}

	var err error

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&user)
	err = dec.Decode(&workspace)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	err = server.DB.AutoMigrate(&user).Error
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, userCreated)
}

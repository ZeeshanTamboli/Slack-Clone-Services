package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ZeeshanTamboli/slack-clone-services/api/responses"

	"github.com/ZeeshanTamboli/slack-clone-services/api/models"
)

type UserOutput struct {
	ID         uint32    `gorm:"primary_key;auto_increment;not null" json:"id"`
	Email      string    `gorm:"size:100;not null;unique" json:"email"`
	Workspaces []string  `gorm:"type:text[];not null" json:"workspaces"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// CreateWorkspace : Create user workspace
func (server *Server) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	userOutput := models.UserOutput{}

	err := json.NewDecoder(r.Body).Decode(&user)
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

	// query := "INSERT INTO users(email, workspaces, created_at, updated_at) VALUES (?, ARRAY[?], ?, ?)"
	// userCreated, err := server.DB.Exec(query, user.Email, user.Workspace, user.CreatedAt, user.UpdatedAt).Error
	// scannedUser := userCreated.Scan(&user.Email)
	// fmt.Println("Scanned user: ", scannedUser)

	err = server.DB.Debug().AutoMigrate(&userOutput).Error
	userCreated, err := userOutput.SaveUser(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, userCreated)
}

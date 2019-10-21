package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// User : Fields of a user
// type User struct {
// 	ID        uint32    `gorm:"primary_key;auto_increment;not null" json:"id"`
// 	Email     string    `gorm:"size:100;not null;unique" json:"email"`
// 	Workspace string    `gorm:"size:100;not null" json:"workspace"`
// 	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// }
type User struct {
	ID        uint32    `json:"id"`
	Email     string    `json:"email"`
	Workspace string    `json:"workspace"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserOutput struct {
	ID         uint32         `gorm:"primary_key;auto_increment;not null" json:"id"`
	Email      string         `gorm:"size:100;not null;unique" json:"email"`
	Workspaces pq.StringArray `gorm:"type:text[];not null" json:"workspaces"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare : Prepares a new User object
func (u *User) Prepare() {
	// u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Workspace = html.EscapeString(strings.TrimSpace(u.Workspace))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate : Validates the given input data
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	if u.Workspace == "" {
		return errors.New("Workspace is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

// SaveUser : Save the user in the database
func (u *UserOutput) SaveUser(db *gorm.DB) (*UserOutput, error) {
	if err := db.Debug().Create(&u).Error; err != nil {
		return &UserOutput{}, err
	}
	return u, nil
}

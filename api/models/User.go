package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

// User : Fields of a user
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment;not null" json:"id"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare : Prepares a new User object
func (u *User) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate : Validates the given input data
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

// SaveUser : Save the user in the database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	if err := db.Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

package models

import (
	"html"
	"strings"
	"time"
)

// User : Fields of a user
type User struct {
	ID         uint32    `gorm:"primary_key;auto_increment;not null" json:"id"`
	Email      string    `gorm:"size:100;not null;unique" json:"email"`
	Workspaces []string  `gorm:"size:100;not null" json:"workspaces"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare : Prepares a new User object
func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	// u.Workspaces =
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

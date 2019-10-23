package models

type Workspace struct {
	ID   uint32 `gorm:"primary_key;auto_increment;not null" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
}

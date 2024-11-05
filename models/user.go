package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:text;not null" json:"name"`
	Email    string `gorm:"type:text;uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:text;not null" json:"password,omitempty"`
}

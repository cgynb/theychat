package model

import (
	"gorm.io/gorm"
)

type User struct {
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"-"`
	Email    string `json:"email"`
	gorm.Model
}

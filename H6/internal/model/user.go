package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"-" validate:"required"`
}

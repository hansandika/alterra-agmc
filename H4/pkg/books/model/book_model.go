package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Author        string `json:"author" validate:"required"`
	YearPublished int    `json:"year_published" validate:"required"`
}

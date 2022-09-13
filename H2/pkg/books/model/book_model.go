package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title         string `json:"title"`
	Description   string `json:"description"`
	Author        string `json:"author"`
	YearPublished int    `json:"year_published"`
}

package dto

type NewBook struct {
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Author        string `json:"author" validate:"required"`
	YearPublished int    `json:"year_published" validate:"required"`
}

type BookResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Author        string `json:"author"`
	YearPublished int    `json:"year_published"`
}

package dto

type NewBook struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Author        string `json:"author"`
	YearPublished int    `json:"year_published"`
}

package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	Title       string           `json:"title" `
	Price       int              `json:"price" `
	Image       string           `json:"image" `
	CategoryId  int              `json:"category_id" `
	Category    CategoryResponse `json:"category"`
	FilmUrl     string           `json:"film_url"`
	Description string           `json:"description" `
	Thumbnail   string           `json:"thumbnail"`
}

type FilmResponse struct {
	gorm.Model
	Title       string           `json:"title" `
	Price       int              `json:"price" `
	Image       string           `json:"image" `
	CategoryId  int              `json:"category_id" `
	Category    CategoryResponse `json:"category"`
	FilmUrl     string           `json:"film_url"`
	Description string           `json:"description" `
	Thumbnail   string           `json:"thumbnail"`
}

func (FilmResponse) TableName() string {
	return "films"
}

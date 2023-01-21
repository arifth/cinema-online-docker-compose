package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type CategoryResponse struct {
	gorm.Model
	Name string `json:"name"`
}

func (CategoryResponse) TableName() string {
	return "categories"
}

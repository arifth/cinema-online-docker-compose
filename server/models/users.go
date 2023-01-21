package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Full_name string `json:"full_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
	// BookId    int          `json:"book_id"`
	// Books     BookResponse `json:"book"`
}

type UserResponse struct {
	gorm.Model
	Full_name string `json:"full_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
	// BookId    int          `json:"book_id"`
	// Books     BookResponse `json:"book"`
}

func (UserResponse) TableName() string {
	return "users"
}

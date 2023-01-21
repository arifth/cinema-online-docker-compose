package repositories

import (
	"final-task/models"

	"github.com/k0kubun/pp/v3"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	GetUserID(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	// func below create user and insert it into db with gorm create method api
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) (models.User, error) {

	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) GetUserID(ID int) (models.User, error) {
	var user models.User

	err := r.db.First(&user, ID).Error

	pp.Println(user)

	return user, err

}

package repositories

import (
	"final-task/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// this func begin handling database items with object relation models
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Debug().Find(&users).Error // Using Find method

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Debug().First(&user, ID).Error // Using First method

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(name,email,password,Created_at,Updated_at) VALUES(?,?,?,?,?)", user.Full_name, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err

}

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Full_name, user.Email, user.Password, ID).Scan(&user).Error

	return user, err
}

// func (r *repository) UpdateProfile(user models.User, ID int) (models.User, error) {
// 	err := r.db.Raw("UPDATE users SET image=? WHERE id=?", user.Image, ID).Scan(&user).Error

// 	return user, err
// }

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

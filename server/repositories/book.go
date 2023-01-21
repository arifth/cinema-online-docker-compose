package repositories

import (
	"final-task/models"

	"github.com/k0kubun/pp/v3"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindBooks() ([]models.Book, error)
	FindBook(id int) (models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	UpdateBook(status string, ID uint) (models.Book, error)
	// DeleteTrip(trip models.Trip, id int) (models.Trip, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBooks() ([]models.Book, error) {
	var book []models.Book
	// err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error

	err := r.db.Debug().Preload("User").Preload("Film.Category").Find(&book).Error
	// err := r.db.Debug().Find(&book).Error
	return book, err
}

// this func begin handling database items with object relation models
func (r *repository) FindBook(id int) (models.Book, error) {
	var book models.Book
	err := r.db.Preload("Film").Preload("Film.Category").First(&book, id).Error // Using Find method

	return book, err
}

func (r *repository) CreateBook(book models.Book) (models.Book, error) {
	// err := r.db.Find("INSERT INTO trips(title,country,accomodation,transportation,eat,day,night,dateTrip,price,quota,description,image) ,
	err := r.db.Debug().Create(&book).Error

	return book, err

}

func (r *repository) UpdateBook(status string, ID uint) (models.Book, error) {

	pp.Println(status)

	var new models.Book
	r.db.Preload("Trip").First(&new, ID)

	pp.Println("after Query", new)

	if status != new.Status && status == "success" {
		var film models.Film
		r.db.First(&film, new.Film.ID)
		// trip.Quota = trip.Quota - new.CounterQty
		r.db.Model(&film).Updates(&film)
	}

	new.Status = status

	pp.Println("after changed status ", new)

	err := r.db.Model(&new).Updates(new).Error
	return new, err
}

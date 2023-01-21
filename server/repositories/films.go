package repositories

import (
	"final-task/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilm() ([]models.Film, error)
	FindSingleFilm(id int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film, id int) (models.Film, error)
	DeleteFilm(film models.Film, id int) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilm() ([]models.Film, error) {
	var film []models.Film
	// err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error

	err := r.db.Preload("Category").Find(&film).Error

	return film, err
}

// this func begin handling database items with object relation models
func (r *repository) FindSingleFilm(id int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category").First(&film, id).Error // Using Find method

	// fmt.Println(trip)

	return film, err
}

func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	// err := r.db.Find("INSERT INTO trips(title,country,accomodation,transportation,eat,day,night,dateTrip,price,quota,description,image) ,
	err := r.db.Preload("Category").Create(&film).Error

	return film, err

}

func (r *repository) UpdateFilm(film models.Film, id int) (models.Film, error) {

	err := r.db.Save(&film).Error

	// err := r.db.Debug().Raw(`"UPDATE trips SET title=?, country_id=?, accomodation=?,transportation=?, eat=?, day=?, night=?, dateTrip=?, price=?, quota=?, description=?, image=? WHERE id=?"`, trip.Title, trip.Country, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	// err := r.db.Debug().Raw("UPDATE trips SET title=?, accomodation=?,transportation=?, eat=?, day=?, night=?, date_trip=?, price=?, quota=?, description=?, image=? WHERE id=?", trip.Title, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	return film, err
}

func (r *repository) DeleteFilm(film models.Film, ID int) (models.Film, error) {

	// err := r.db.Preload("Country").Delete(&trip).Error
	err := r.db.Raw("DELETE FROM films WHERE id=?", ID).Scan(&film).Error
	return film, err
}

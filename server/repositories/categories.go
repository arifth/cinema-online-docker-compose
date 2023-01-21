package repositories

import (
	"final-task/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategories() ([]models.Category, error)
	FindCategory(id int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category, id int) (models.Category, error)
	DeleteCategory(category models.Category, id int) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategories() ([]models.Category, error) {
	var category []models.Category
	// err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error

	err := r.db.Find(&category).Error

	return category, err
}

// this func begin handling database items with object relation models
func (r *repository) FindCategory(id int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error // Using Find method

	// fmt.Println(category)

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	// err := r.db.Find("INSERT INTO trips(title,country,accomodation,transportation,eat,day,night,dateTrip,price,quota,description,image) ,
	err := r.db.Debug().Preload("Category").Create(&category).Error

	return category, err

}

func (r *repository) UpdateCategory(category models.Category, id int) (models.Category, error) {

	// NOTES: sementara untuk supy tidak error
	// err := r.db.Save(&country).Error
	err := r.db.Raw("UPDATE categories SET name=? WHERE id=?", category.Name, id).Scan(&category).Error

	// err := r.db.Debug().Raw(`"UPDATE trips SET title=?, country_id=?, accomodation=?,transportation=?, eat=?, day=?, night=?, dateTrip=?, price=?, quota=?, description=?, image=? WHERE id=?"`, trip.Title, trip.Country, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	// err := r.db.Debug().Raw("UPDATE trips SET title=?, accomodation=?,transportation=?, eat=?, day=?, night=?, date_trip=?, price=?, quota=?, description=?, image=? WHERE id=?", trip.Title, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	return category, err
}

func (r *repository) DeleteCategory(category models.Category, ID int) (models.Category, error) {

	// err := r.db.Preload("Country").Delete(&trip).Error
	err := r.db.Raw("DELETE FROM categories WHERE id=?", ID).Scan(&category).Error
	return category, err
}

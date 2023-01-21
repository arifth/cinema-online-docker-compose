package database

import (
	"final-task/models"
	"final-task/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {

	// NOTE: migrate parent models first ,then child models after, otherwise it wont create the relation table between both

	err := mysql.DB.AutoMigrate(&models.User{}, &models.Film{}, &models.Book{}, &models.Book{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Succes")
}

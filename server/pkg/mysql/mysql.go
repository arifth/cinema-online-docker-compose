package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	// dsn := "root:MIQkFRszq3Ichol04MnX@tcp(containers-us-west-24.railway.app:6395)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(golang-docker-db:3306)/final-task?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}

package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:pass1234@tcp(localhost:3306)/rest-api-go"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}

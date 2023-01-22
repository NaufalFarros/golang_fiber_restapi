package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost)/go_restapi_fiber"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&book{})
	DB = db
}

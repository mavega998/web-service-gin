package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=postgres password=123456 dbname=postgres sslmode=disable"
var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Album{})
	if err != nil {
		return
	}

	DB = database
}

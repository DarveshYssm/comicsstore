package config

import (
	"comics-store/models"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=Darveshova2001 dbname=comics_store port=5433 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database.AutoMigrate(&models.Author{}, &models.Category{}, &models.Comic{}, &models.User{})
	DB = database
}

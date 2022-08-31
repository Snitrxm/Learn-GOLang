package database

import (
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"projeto/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb(){
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB\n", err.Error())
		os.Exit(2)
	}

	log.Println("Database connect")
	log.Println("Running Migrations")
	// Add Migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
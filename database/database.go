package database

import (
	"log"
	"os"

	"github.com/Radser2001/products_api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	//check if there is an error
	if err != nil {
		log.Fatal("Failed to connect to the database")
		os.Exit(2)
	}

	//if there are no errors
	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	println("Running Migrations")

	//add migrations
	db.AutoMigrate(&models.Product{})

	Database = DbInstance{Db: db}

}

package database

import (
	"log"
	"os"

	"github.com/chmenegatti/gofiber-api/models"
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

	if err != nil {
		log.Fatal("Failed to connect do the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully!")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	//TODO: add migrations

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}

}

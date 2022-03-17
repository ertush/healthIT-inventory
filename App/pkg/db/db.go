package db

import (
	"log"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=127.0.0.1 user=pg password=pass dbname=crud port=5432 sslmode=disable "

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Auto migrate DB
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Categories{})
	db.AutoMigrate(&models.Dispatch{})
	db.AutoMigrate(&models.Equipment{})

	return db
}

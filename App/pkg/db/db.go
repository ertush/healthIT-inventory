package db

import (
	"fmt"
	"log"

	"github.com/ertush/healthIT-inventory/App/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dbCreds := map[string]string{
		"user":     "pg",
		"password": "pass",
		"dbname":   "crud",
		"host":     "db",
		"port":     "5432",
	}

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbCreds["user"],
		dbCreds["password"],
		dbCreds["host"],
		dbCreds["port"],
		dbCreds["dbname"])

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	// Check if DB has connected
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

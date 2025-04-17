package db

import (
	"log"
	"os"
	"test-ara/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
func InitDB() {
	godotenv.Load()

	var err error
	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	db.AutoMigrate(&models.Film{}, &models.ProductionHouse{})
}

func GetDBInstance() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}


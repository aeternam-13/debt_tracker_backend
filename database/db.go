package database

import (
	"debt_tracker/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("debttracker.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	db.AutoMigrate(&models.User{})
}

package db

import (
	"github.com/kittichok/event-driven/payment/src/db/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase(d gorm.Dialector) {
	database, err := gorm.Open(d, &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Payment{})

	DB = database
}

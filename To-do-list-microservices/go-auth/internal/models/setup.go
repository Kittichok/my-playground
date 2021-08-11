package models

import (
	"github.com/kittichok/go-auth/internal/utils"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase(d gorm.Dialector) {
	database, err := gorm.Open(d, &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Token{})

	DB = database
}

func SeedUser(u User) {
	salt, err := utils.GenerateRandomBytes(utils.SaltSize)
	if err != nil {
		panic("Failed to generate salt")
	}
	hash := utils.HashPassword([]byte(u.Password), salt)
	u.Password = string(hash)
	u.Salt = string(salt)
	DB.Create(u)
}

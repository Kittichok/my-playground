package main

import (
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/server"
	"gorm.io/driver/sqlite"
)


func main() {
	d := sqlite.Open("test.db")
	models.ConnectDataBase(d)
	
	server.Init()
}


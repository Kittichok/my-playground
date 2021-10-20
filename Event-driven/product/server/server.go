package server

import (
	"gorm.io/driver/sqlite"

	"github.com/kittichok/event-driven/product/controllers"
	"github.com/kittichok/event-driven/product/db"
	"github.com/kittichok/event-driven/product/repository"
	"github.com/kittichok/event-driven/product/usecase"
)

func Init() {
	d := sqlite.Open("test.db")
	db.ConnectDataBase(d)
	rep := repository.NewUserRepository(db.DB)
	u := usecase.NewProductUseCase(rep)
	c := controllers.NewController(u)
	r := SetupRouter(c)

	//TODO use viper get env
	port := 5000
	r.Run(":" + port)
}

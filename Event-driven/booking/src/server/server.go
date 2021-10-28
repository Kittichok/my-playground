package server

import (
	"github.com/kittichok/event-driven/booking/src/controllers"
	"github.com/kittichok/event-driven/booking/src/db"
	"github.com/kittichok/event-driven/booking/src/db/repository"
	"github.com/kittichok/event-driven/booking/src/event"
	"github.com/kittichok/event-driven/booking/src/usecase"
	"gorm.io/driver/sqlite"
)

func Init() {
	d := sqlite.Open("test.db")
	db.ConnectDataBase(d)
	rep := repository.NewRepository(db.DB)
	e := event.NewEventConnection()
	u := usecase.NewUseCase(rep, e)
	c := controllers.NewController(u)
	r := SetupRouter(c)

	go event.NewConsumer()

	//TODO use viper get env
	port := "4000"
	r.Run(":" + port)
}

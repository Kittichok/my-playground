package server

import (
	"github.com/kittichok/event-driven/payment/src/controllers"
	"github.com/kittichok/event-driven/payment/src/db"
	"github.com/kittichok/event-driven/payment/src/db/repository"
	"github.com/kittichok/event-driven/payment/src/event"
	"github.com/kittichok/event-driven/payment/src/event/processor"
	"github.com/kittichok/event-driven/payment/src/usecase"
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

	go processor.NewConsumer(u)

	//TODO use viper get env
	port := "4001"
	r.Run(":" + port)
}

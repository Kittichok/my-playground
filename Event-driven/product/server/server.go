package server

import (
	"go.uber.org/dig"

	"github.com/kittichok/event-driven/product/controllers"
	"github.com/kittichok/event-driven/product/db"
	"github.com/kittichok/event-driven/product/db/repository"
	"github.com/kittichok/event-driven/product/usecase"
)

type ServerService struct {
	dig.In

	Controller controllers.IController
}

func Init() {
	container := dig.New()
	container.Provide(db.NewSqliteConnection)
	container.Provide(repository.NewProductRepository)
	container.Provide(usecase.NewProductUseCase)
	container.Provide(controllers.NewController)

	httpServ := func(s ServerService) {
		r := SetupRouter(s.Controller)
		port := "5001"
		r.Run(":" + port)
	}
	if err := container.Invoke(httpServ); err != nil {
		panic(err)
	}
}

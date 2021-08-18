package server

import (
	"os"

	"github.com/kittichok/go-auth/internal/controllers"
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/internal/repository"
	"github.com/kittichok/go-auth/internal/usecase/users"
	"gorm.io/driver/sqlite"
)

func Init() {
	d := sqlite.Open("test.db")
	models.ConnectDataBase(d)
	rep := repository.NewUserRepository(models.DB)
	u := users.NewUserUseCase(rep)
	authController := controllers.NewAuthController(u)
	r := SetupRouter(authController)

	port := getenv("PORT", "80")
	r.Run(":" + port)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
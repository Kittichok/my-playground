package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kittichok/app/internal/controllers"
	"github.com/kittichok/app/internal/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	// models.Seed()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1/api")
	{
		v1.POST("/signin", controllers.SignIn)
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/tokens", controllers.GetTokens)
		v1.POST("/signup", controllers.SignUp)
		//TODO add api refresh token
	}
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

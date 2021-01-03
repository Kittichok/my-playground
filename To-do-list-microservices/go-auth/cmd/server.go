package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kittichok/app/internal/controllers"
	"github.com/kittichok/app/internal/models"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

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

	v1 := r.Group("/v1")
	{
		v1.POST("/api/signin", controllers.SignIn)
		v1.GET("/api/users", controllers.GetUsers)
		v1.GET("/api/tokens", controllers.GetTokens)
		//TODO add api refresh token
		//TODO add api create user
	}

	r.Run(":" + port)
}

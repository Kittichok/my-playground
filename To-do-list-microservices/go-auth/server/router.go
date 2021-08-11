package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittichok/go-auth/internal/controllers"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
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
	return r
}

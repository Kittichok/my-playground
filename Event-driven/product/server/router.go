package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittichok/event-driven/product/controllers"
)

func SetupRouter(c controllers.AuthController) *gin.Engine {
	r := gin.Default()

	// r.Use(sentrygin.New(sentrygin.Options{}))
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/products", c.GetProductList)
	}
	return r
}

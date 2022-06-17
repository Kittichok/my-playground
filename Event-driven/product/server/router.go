package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittichok/event-driven/product/controllers"
)

func SetupRouter(c controllers.IController) *gin.Engine {
	r := gin.Default()

	// r.Use(sentrygin.New(sentrygin.Options{}))
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"*"}

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(c.Tracer)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/products", c.GetProductList)
		v1.POST("/product", c.AddProduct)
	}
	return r
}

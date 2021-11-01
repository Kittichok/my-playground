package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittichok/event-driven/booking/src/controllers"
)

func SetupRouter(c controllers.Controller) *gin.Engine {
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
		v1.POST("/booking", c.CreateBooking)
		v1.PATCH("/booking", c.UpdateBooking)
		v1.GET("/booking/:bookingID/submit", c.SubmitBooking)
	}
	return r
}

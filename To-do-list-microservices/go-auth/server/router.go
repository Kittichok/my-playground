package server

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kittichok/go-auth/internal/controllers"
)


func SetupRouter(c controllers.AuthController) *gin.Engine {
	r := gin.Default()

	r.Use(sentrygin.New(sentrygin.Options{}))
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
		v1.POST("/signin", c.SignIn)
		v1.GET("/users", c.GetUsers)
		// v1.GET("/tokens", c.GetTokens)
		v1.POST("/signup", c.SignUp)
		//TODO add api refresh token
	}
	return r
}

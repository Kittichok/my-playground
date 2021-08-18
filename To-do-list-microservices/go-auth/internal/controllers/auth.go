package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/internal/pkg"
	"github.com/kittichok/go-auth/internal/usecase/users"
)

type AuthController interface {
	SignIn(*gin.Context)
	SignUp(*gin.Context)
	GetUsers(*gin.Context)
	GetTokens(*gin.Context)
}

type authController struct {
	usecase users.UseCase
}

func NewAuthController(u users.UseCase) AuthController {
	return authController{
		usecase: u,
	}
}

func (a authController) SignIn(c *gin.Context) {
	span := pkg.StartSpan(c.Request.Context(), "signin")
	var json *models.User
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := a.usecase.SignIn(*json)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pkg.FinishSpan(span)
	c.JSON(200, r)
	return
}

func (a authController) GetUsers(c *gin.Context) {
	span := pkg.StartSpan(c.Request.Context(), "get users")
	var users []models.User
	err := models.DB.Find(&users).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pkg.FinishSpan(span)
	c.JSON(200, gin.H{
		"users": users,
	})
	return
}

func (a authController) GetTokens(c *gin.Context) {
	span := pkg.StartSpan(c.Request.Context(), "get tokens")
	var tokens []models.Token
	err := models.DB.Find(&tokens).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pkg.FinishSpan(span)
	c.JSON(200, gin.H{
		"tokens": tokens,
	})
	return
}

func (a authController) SignUp(c *gin.Context) {
	span := pkg.StartSpan(c.Request.Context(), "signup")
	var json *models.User
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.usecase.SignUp(*json)
	if err != nil {
		pkg.FinishSpan(span)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pkg.FinishSpan(span)
	c.Status(http.StatusCreated)
	return
}

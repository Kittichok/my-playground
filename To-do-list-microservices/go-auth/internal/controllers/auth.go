package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kittichok/app/internal/models"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var json *models.User
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existUser models.User
	/// @TODO password must hash
	err = models.DB.First(&existUser, json).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	currentTime := time.Now().AddDate(0, 0, 2).Unix()
	expiresAt := int64(currentTime)
	claims := &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Issuer:    "auth.todo-list",
		Id:        strconv.FormatUint(uint64(existUser.ID), 10),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	mySigningKey := []byte("AllYourBase")
	tokenString, _ := accessToken.SignedString(mySigningKey)

	t := models.Token{
		AccessToken: tokenString,
		IsActive:    true,
		UserID:      existUser.ID,
	}
	// err = existUser.AddToken(t)
	err = models.SaveToken(t)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"AccessToken": tokenString,
		"TokenType":   "Bearer",
		"ExpiresIn":   expiresAt,
	})
	return
}

func GetUsers(c *gin.Context) {
	var users []models.User
	err := models.DB.Find(&users).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"users": users,
	})
	return
}

func GetTokens(c *gin.Context) {
	var tokens []models.Token
	err := models.DB.Find(&tokens).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"tokens": tokens,
	})
	return
}

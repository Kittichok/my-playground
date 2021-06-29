package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/internal/utils"

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
	// @TODO : refactor
	err = models.DB.First(&existUser, models.User{Username: json.Username}).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}
	json.Password = string(utils.HashPassword([]byte(json.Password), []byte(existUser.Salt)))
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
		// Id:        strconv.FormatUint(uint64(existUser.ID), 10),
		Id: existUser.ID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	mySigningKey := []byte("AllYourBase")
	tokenString, _ := accessToken.SignedString(mySigningKey)

	t := models.Token{
		AccessToken: tokenString,
		IsActive:    true,
		UserID:      existUser.ID,
	}
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

func SignUp(c *gin.Context) {
	var json *models.User
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existUser models.User
	models.DB.First(&existUser, models.User{Username: json.Username})
	if existUser.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is exist"})
		return
	}

	//TODO change store to pass.salt format??
	salt, err := utils.GenerateRandomBytes(utils.SaltSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "system error"})
		return
	}
	hash := utils.HashPassword([]byte(json.Password), salt)
	json.Password = string(hash)
	json.Salt = string(salt)
	json.ID = uuid.Must(uuid.NewRandom()).String()
	var user = models.DB.Create(&json)
	// var user = models.DB.Create(models.User{Username: json.Username, Password: json.Password, Salt: json.Salt})
	if user != nil {
		c.Status(http.StatusCreated)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "system error"})
	return
}

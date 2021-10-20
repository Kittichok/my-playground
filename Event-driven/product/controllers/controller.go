package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittichok/event-driven/product/models"
)

type IController interface {
	GetProductList(*gin.Context)
}

type Controller struct {
	usecase UseCase
}

func NewAuthController(u product.UseCase) Controller {
	return Controller{
		usecase: u,
	}
}

func (c Controller) GetProductList(ctx *gin.Context) {
	var products []models.product
	err := models.DB.Find(&products).Error
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"product": products,
	})
	return
}

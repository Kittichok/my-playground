package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittichok/event-driven/product/db/models"
	"github.com/kittichok/event-driven/product/usecase"
)

type IController interface {
	GetProductList(*gin.Context)
	AddProduct(*gin.Context)
}

type Controller struct {
	usecase usecase.IUseCase
}

func NewController(u usecase.IUseCase) IController {
	return Controller{
		usecase: u,
	}
}

func (c Controller) GetProductList(ctx *gin.Context) {
	products, err := c.usecase.GetProductList()
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

func (c Controller) AddProduct(ctx *gin.Context) {

	var json models.Product
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.usecase.AddProduct(json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
	return
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kittichok/event-driven/payment/src/usecase"
)

type Controller struct {
	usecase usecase.IUseCase
}

func NewController(usecase usecase.IUseCase) Controller {
	return Controller{usecase}
}

func (c *Controller) CreatePaymentMethod(ctx *gin.Context) {
	// var json usecase.ReqCreateBooking
	// err := ctx.ShouldBindJSON(&json)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// err = c.usecase.CreatePaymentMethod(json)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	ctx.Status(http.StatusForbidden)
}

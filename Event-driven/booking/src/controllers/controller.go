package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kittichok/event-driven/booking/src/usecase"
)

type Controller struct {
	usecase usecase.IUseCase
}

func NewController(usecase usecase.IUseCase) Controller {
	return Controller{usecase}
}

func (c Controller) CreateBooking(ctx *gin.Context) {
	var json usecase.ReqCreateBooking
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respData, err := c.usecase.CreateBooking(json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, respData)
}

func (c Controller) UpdateBooking(ctx *gin.Context) {
	var json usecase.ReqUpdateBooking
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.usecase.UpdateBooking(json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c Controller) SubmitBooking(ctx *gin.Context) {
	id := ctx.Param("bookingID")
	bookingID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "booking id not found"})
		return
	}

	err = c.usecase.SubmitBooking(ctx, bookingID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusAccepted)
}

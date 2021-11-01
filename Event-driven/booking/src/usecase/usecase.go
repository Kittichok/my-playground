package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kittichok/event-driven/booking/src/db/models"
	"github.com/kittichok/event-driven/booking/src/db/repository"
	"github.com/kittichok/event-driven/booking/src/event"
)

type Product struct {
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}
type ReqCreateBooking struct {
	UserID   int64     `json:"user_id" required:"true"`
	Products []Product `json:"products" required:"true"`
}

type ReqUpdateBooking struct {
	BookID   int64     `json:"book_id" required:"true"`
	Products []Product `json:"products" required:"true"`
}

type IUseCase interface {
	CreateBooking(req ReqCreateBooking) error
	UpdateBooking(req ReqUpdateBooking) error
	SubmitBooking(ctx context.Context, bookingID int64) error
}

type UseCase struct {
	repo  repository.IRepository
	event event.Event
}

func NewUseCase(repo repository.IRepository, event event.Event) IUseCase {
	return UseCase{repo, event}
}

func (c UseCase) CreateBooking(req ReqCreateBooking) error {
	booking, err := c.repo.AddBooking(models.NewBooking(req.UserID))
	if err != nil {
		fmt.Errorf("Error: %v\n", err)
		return err
	}
	var details []models.BookingDetail
	for _, p := range req.Products {
		details = append(details, models.NewBookingDetail(
			booking.ID,
			p.ProductID,
			p.Quantity,
			p.Price,
		))
	}
	err = c.repo.AddBookingDetail(details)
	if err != nil {
		return err
	}
	return nil
}
func (c UseCase) UpdateBooking(req ReqUpdateBooking) error {
	return nil
}

func (c UseCase) SubmitBooking(ctx context.Context, id int64) error {
	condition := &models.Booking{ID: id}
	booking, err := c.repo.FindBooking(*condition)
	if err != nil {
		return err
	}
	bookingDetails, err := c.repo.FindAllBookingDetail(models.BookingDetail{BookID: id})
	if err != nil {
		return err
	}
	msg := event.BookingSubmitBody{
		Booking:       *booking,
		BookingDetail: bookingDetails,
	}
	body, err := json.Marshal(&msg)
	if err != nil {
		return err
	}
	c.event.SubmitMessage(ctx, event.BookingSubmit, string(body))
	// TODO submit book and sent event payment to payment service
	// and update product stock to product service

	return nil
}

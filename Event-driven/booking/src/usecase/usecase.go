package usecase

import (
	"github.com/kittichok/event-driven/booking/src/db/repository"
	"github.com/kittichok/event-driven/booking/src/event"
)

type ReqCreateBooking struct {
}

type ReqUpdateBooking struct {
}

type IUseCase interface {
	CreateBooking(req ReqCreateBooking) error
	UpdateBooking(req ReqUpdateBooking) error
}

type UseCase struct {
	repo  repository.IRepository
	event event.Event
}

func NewUseCase(repo repository.IRepository, event event.Event) IUseCase {
	return UseCase{repo, event}
}

func (c UseCase) CreateBooking(req ReqCreateBooking) error {
	// c.rep.
	return nil
}
func (c UseCase) UpdateBooking(req ReqUpdateBooking) error {
	return nil
}

func (c UseCase) SubmitBooking(id int64) error {
	// details := c.repo.FindMany(models.BookingDetail, id)
	// c.event.payment
	//TODO submit book and sent event payment to payment service
	// and update product stock to product service
	return nil
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingDetail struct {
	gorm.Model
	ID        int64 `gorm:"primary_key"`
	BookID    int64
	ProductID int64
	Quantity  int64
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBookingDetail(bookID int64, productID int64, quantity int64, price float64) BookingDetail {
	return BookingDetail{
		BookID:    bookID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

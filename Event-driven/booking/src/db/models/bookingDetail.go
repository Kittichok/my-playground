package models

import "gorm.io/gorm"

type BookingDetail struct {
	gorm.Model
	ID        int64 `gorm:"primary_key"`
	BookID    int64
	ProductID int64
	Quantity  int64
	Price     float64
}

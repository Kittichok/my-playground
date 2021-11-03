package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentStatus string

const (
	Pending   PaymentStatus = "pending"
	Paid      PaymentStatus = "paid"
	OnProcess PaymentStatus = "on process"
)

type Payment struct {
	gorm.Model
	ID            int64 `gorm:"primary_key"`
	OwnerID       int64
	PaymentStatus PaymentStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewPayment(ownerID int64) Payment {
	return Payment{
		OwnerID:       ownerID,
		PaymentStatus: Pending,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
}

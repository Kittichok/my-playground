package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID      int64 `gorm:"primary_key"`
	OwnerID int64
}

package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       int64   `gorm:"primary_key"`
	Name     string  `binding:"required"`
	Quantity int32   `binding:"required"`
	Price    float64 `binding:"required"`
	Active   bool
}

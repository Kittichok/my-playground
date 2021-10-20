package models

type Product struct {
	ID       int64   `gorm:"primary_key"`
	Name     string  `binding:"required" gorm:"primary_key"`
	Quantity int32   `binding:"required"`
	Price    float64 `binding:"required"`
	Active   bool
}

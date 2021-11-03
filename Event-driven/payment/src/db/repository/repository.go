package repository

import (
	"gorm.io/gorm"
)

type IRepository interface {
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return Repository{
		DB: db,
	}
}

package repository

import (
	"github.com/kittichok/event-driven/product/models"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Find(models.Product) (*models.Product, error)
	Add(p models.Product) error
}

type productRepository struct {
	DB gorm.DB
}

func NewProductRepository(DB *gorm.DB) IProductRepository {
	return productRepository{DB}
}

func (pRep productRepository) Add(p models.Product) error {
	return nil
}

func (pRep productRepository) Find(p models.Product) (*models.Product, error) {
	return nil, nil
}

package repository

import (
	"github.com/kittichok/event-driven/product/db/models"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Find(p models.Product) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Add(p models.Product) error
	Update(p models.Product) error
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) IProductRepository {
	return productRepository{DB: DB}
}

func (pRep productRepository) Add(p models.Product) error {
	result := pRep.DB.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pRep productRepository) Find(p models.Product) (*models.Product, error) {
	var product models.Product
	result := pRep.DB.First(&product, p)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (pRep productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	result := pRep.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (pRep productRepository) Update(p models.Product) error {
	result := pRep.DB.Save(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

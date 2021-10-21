package usecase

import (
	"github.com/kittichok/event-driven/product/db/models"
	"github.com/kittichok/event-driven/product/db/repository"
)

type RespProduct struct {
	Name     string
	Quantity int32
	Price    float64
}

type IUseCase interface {
	GetProductList() ([]models.Product, error)
	// UpdateProduct(models.Product) error
	AddProduct(models.Product) error
	// InactiveProduct(int64) error
}

type ProductUsecase struct {
	rep repository.IProductRepository
}

func NewProductUseCase(rep repository.IProductRepository) IUseCase {
	return ProductUsecase{rep}
}

func (p ProductUsecase) GetProductList() ([]models.Product, error) {
	products, err := p.rep.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductUsecase) AddProduct(product models.Product) error {
	err := p.rep.Add(product)
	if err != nil {
		return err
	}
	return nil
}

package usecase

import (
	"github.com/kittichok/event-driven/product/models"
	"github.com/kittichok/event-driven/product/repository"
)

type RespProduct struct {
	Name     string
	Quantity int32
	Price    float64
}

type IUseCase interface {
	GetProductList() ([]*RespProduct, error)
	UpdateProduct(models.Product) error
	AddProduct(models.Product) error
	InactiveProduct(int64) error
}

type ProductUsecase struct {
	rep repository.IProductRepository
}

func NewProductUseCase(rep repository.IProductRepository) ProductUsecase {
	return ProductUsecase{rep}
}

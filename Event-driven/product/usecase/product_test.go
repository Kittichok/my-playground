package usecase

import (
	"testing"

	"github.com/kittichok/event-driven/product/db/models"
	"github.com/kittichok/event-driven/product/db/repository"
)

func TestAddProduct(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	usecase := NewProductUseCase(repo)

	product := models.Product{
		Name:     "test",
		Quantity: 1,
		Price:    1.0,
	}

	usecase.AddProduct(product)

	products, _ := usecase.GetProductList()
	if len(products) < 1 {
		t.Errorf("Product not found")
	}
}

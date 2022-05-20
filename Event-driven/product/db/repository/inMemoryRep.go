package repository

import (
	"github.com/kittichok/event-driven/product/db/models"
)

type InMemoryRepository struct {
	db []models.Product
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		db: make([]models.Product, 0),
	}
}

func (r *InMemoryRepository) Find(p models.Product) (*models.Product, error) {
	return &p, nil
}

func (r *InMemoryRepository) FindAll() ([]models.Product, error) {
	return r.db, nil
}
func (r *InMemoryRepository) Add(p models.Product) error {
	r.db = append(r.db, p)
	return nil
}
func (r *InMemoryRepository) Update(p models.Product) error {
	return nil
}

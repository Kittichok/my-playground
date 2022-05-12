package repository

type IRepository[T any] interface {
	// Find(id string) (T, error)
	FindByNameAndPrice(name string, price float64) ([]T, error)
	// Insert(T) error
	// Update(T) error
	// Delete(id string) error
}

type Repository[T any] struct {
	DB []T
}

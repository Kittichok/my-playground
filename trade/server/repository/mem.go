package repository

func NewMemRepository[T any]() Repository[T] {
	return Repository[T]{
		DB: []T{},
	}
}

func (repo *Repository[T]) FindByNameAndPrice(name string, price float64) ([]T, error) {
	// for _, t := range repo.DB {
	// 	if t.Name == name && t.Price == price {
	// 		return []T{t}, nil
	// 	}
	// }
	return []T{}, nil
}

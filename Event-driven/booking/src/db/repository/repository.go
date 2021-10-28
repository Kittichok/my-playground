package repository

import (
	"gorm.io/gorm"
)

type IRepository interface {
	Add(interface{}) error
	Update(interface{}) error
	Find(interface{}) (interface{}, error)
	FindAll([]interface{}) ([]interface{}, error)
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return Repository{
		DB: db,
	}
}

func (repo Repository) Add(model interface{}) error {
	result := repo.DB.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo Repository) Update(interface{}) error {
	return nil
}

func (repo Repository) Find(interface{}) (interface{}, error) {
	return nil, nil
}

func (repo Repository) FindAll(model []interface{}) ([]interface{}, error) {
	result := repo.DB.Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

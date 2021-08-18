package repository

import (
	"github.com/kittichok/go-auth/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Find(models.User) (*models.User, error)
	FindByName(name string) (*models.User)
	FindById(id string) (*models.User)
	Add(user models.User) (error)
	AddToken(token models.Token) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) IUserRepository {
	return userRepository{DB}
}

func (uRep userRepository) Add(user models.User) (error) {
	var createdUser = uRep.DB.Create(&user)
	if createdUser != nil {
		return nil
	}
	return models.ErrCannotCreate
}

func (uRep userRepository) Find(user models.User) (*models.User, error) {
	var existUser models.User
	err := uRep.DB.First(&existUser, user).Error
	if err != nil {
		return nil, err
	}
	return &existUser, nil
}

func (uRep userRepository) FindByName(username string) (*models.User) {
	var existUser models.User
	err := uRep.DB.First(&existUser, models.User{Username: username}).Error
	if err != nil {
		return nil
	}
	return &existUser
}

func (uRep userRepository) FindById(id string) (*models.User) {
	var existUser models.User
	uRep.DB.First(&existUser, models.User{ID: id})
	if existUser.ID != "" {
		return &existUser
	}
	return nil
}


func (uRep userRepository) AddToken(token models.Token) error {
	err := uRep.DB.Model(&token).Create(token).Error
	if err != nil {
		return err
	}
	return nil
}
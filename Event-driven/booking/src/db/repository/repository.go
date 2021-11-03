package repository

import (
	"log"

	"github.com/kittichok/event-driven/booking/src/db/models"
	"gorm.io/gorm"
)

type IRepository interface {
	Add(interface{}) (interface{}, error)
	AddBooking(models.Booking) (*models.Booking, error)
	AddBookingDetail([]models.BookingDetail) error
	Update(interface{}) error
	UpdateBooking(models.Booking) error
	Find(interface{}) (interface{}, error)
	FindBooking(models.Booking) (*models.Booking, error)
	FindAll(interface{}) (interface{}, error)
	FindAllBookingDetail(models.BookingDetail) ([]models.BookingDetail, error)
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return Repository{
		DB: db,
	}
}

func (repo Repository) Add(model interface{}) (interface{}, error) {
	result := repo.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (repo Repository) AddBooking(b models.Booking) (*models.Booking, error) {
	result := repo.DB.Create(&b)
	if result.Error != nil {
		return nil, result.Error
	}
	return &b, nil
}

func (repo Repository) AddBookingDetail(b []models.BookingDetail) error {
	result := repo.DB.Create(&b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo Repository) Update(interface{}) error {
	return nil
}

func (repo Repository) Find(model interface{}) (interface{}, error) {
	result := repo.DB.Find(&model, model)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (repo Repository) FindAll(model interface{}) (interface{}, error) {
	var results []interface{}
	result := repo.DB.Find(&results, model)
	if result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}

func (repo Repository) FindAllBookingDetail(b models.BookingDetail) ([]models.BookingDetail, error) {
	var results []models.BookingDetail
	result := repo.DB.Find(&results, models.BookingDetail{BookID: b.BookID})
	if result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}

func (repo Repository) FindBooking(b models.Booking) (*models.Booking, error) {
	result := repo.DB.First(&b, models.Booking{ID: b.ID})
	if result.Error != nil {
		return nil, result.Error
	}
	return &b, nil
}

func (repo Repository) UpdateBooking(b models.Booking) error {
	var booking models.Booking
	err := repo.DB.First(&booking, models.Booking{ID: b.ID})

	if err != nil {
		return err.Error
	}
	log.Printf("booking : %v", booking.PaymentStatus)
	log.Printf("b status : %v", b.PaymentStatus)

	booking.PaymentStatus = b.PaymentStatus

	err = repo.DB.Save(&booking)
	if err != nil {
		return err.Error
	}
	return nil
}

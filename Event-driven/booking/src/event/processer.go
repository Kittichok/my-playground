package event

import (
	"encoding/json"
	"fmt"

	"github.com/kittichok/event-driven/booking/src/db/models"
	"github.com/kittichok/event-driven/booking/src/db/repository"
)

func updateBooking(msg string, repo repository.IRepository) {
	b := []byte(msg)
	var booking models.Booking
	err := json.Unmarshal(b, &booking)
	if err != nil {
		fmt.Errorf("event update booking error: %v", err.Error())
	}
	err = repo.UpdateBooking(booking)
	if err != nil {
		fmt.Errorf("event update booking error: %v", err.Error())
	}
}

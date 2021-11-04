package event

import (
	"testing"

	"github.com/kittichok/event-driven/booking/src/db"
	"github.com/kittichok/event-driven/booking/src/db/models"
	"github.com/kittichok/event-driven/booking/src/db/repository"
	"gorm.io/driver/sqlite"
)

// func TestSubmitEvent(t *testing.T) {

// 	e := NewEventConnection()
// 	ctx := context.Background()
// 	e.SubmitMessage(ctx, "BookingSubmit", "{detail: [{pid:1,q:2}]}")

// }

func TestUpdateBooking(t *testing.T) {
	msg := "{\"ID\":1,\"OwnerID\":9999,\"PaymentStatus\":\"paid\"}"
	d := sqlite.Open("file::memory:?cache=shared")
	db.ConnectDataBase(d)
	rep := repository.NewRepository(db.DB)
	rep.AddBooking(models.NewBooking(9999))
	updateBooking(msg, rep)
	t.Fail()
}

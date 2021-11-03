package models

type Booking struct {
	ID            int64
	OwnerID       int64
	PaymentStatus string
}

type BookingDetail struct {
	ID        int64
	BookID    int64
	ProductID int64
	Quantity  int64
	Price     float64
}

type BookingSubmitBody struct {
	Booking       Booking
	BookingDetail []BookingDetail
}

package bookingdomain

// BookingSeat joins Table for Bookings and Seats.
// A specific seat can be booked only once per booking
type BookingSeat struct {
	ID        uint    `json:"id"`
	BookingID uint    `json:"booking_id"`
	SeatID    uint    `json:"seat_id"`
	Price     float64 `json:"price"`
}

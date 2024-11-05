package bookingdomain

import "time"

type Booking struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	ShowtimeID  uint      `json:"showtime_id"`
	BookingTime time.Time `json:"booking_time"`
	TotalPrice  float64   `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Active      bool      `json:"active"`
}

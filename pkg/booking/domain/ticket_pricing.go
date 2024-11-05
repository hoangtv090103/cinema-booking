package bookingdomain

import "time"

// TicketPricing Each seat type has unique pricing per day of the week
type TicketPricing struct {
	ID uint `json:"id"`
	// e.g., "Regular", "VIP"
	SeatType  string    `json:"seat_type"`
	DayOfWeek uint      `json:"day_of_week"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
}

package bookingdomain

import "time"

type Payment struct {
	ID          uint      `json:"id"`
	BookingID   uint      `json:"booking_id"`
	Amount      float64   `json:"amount"`
	PaymentTime time.Time `json:"payment_time"`
	// e.g., "Completed", "Pending", "Failed"
	PaymentStatus string    `json:"payment_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Active        bool      `json:"active"`
}

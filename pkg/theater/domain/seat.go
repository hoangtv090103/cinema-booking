package theaterdomain

import "time"

type Seat struct {
	ID        uint      `json:"id"`
	ScreenID  uint      `json:"screen_id"`
	Row       string    `json:"row"`
	Number    uint      `json:"seat_number"`
	SeatType  string    `json:"seat_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
	// Ensure seat numbers are unique within a screen

}

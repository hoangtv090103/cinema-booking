package bookingdomain

import "time"

type Showtime struct {
	ID        uint      `json:"id"`
	MovieID   uint      `json:"movie_id"`
	ScreenID  uint      `json:"screen_id"`
	StartTime time.Time `json:"start_time"`
	TheaterID uint      `json:"theater_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
}

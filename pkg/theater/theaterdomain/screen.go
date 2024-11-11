package theaterdomain

import "time"

type Screen struct {
	ID         uint      `json:"id"`
	TheaterID  uint      `json:"theater_id"`
	ScreenName string    `json:"name"`
	Capacity   uint      `json:"capacity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Active     bool      `json:"active"`
}

type ScreenCreate struct {
	TheaterID  uint   `json:"theater_id"`
	ScreenName string `json:"name"`
	Capacity   uint   `json:"capacity"`
}

package bookingdomain

type Pricing struct {
	ID        uint    `json:"id"`
	DayOfWeek uint    `json:"day_of_week"`
	SeatType  string  `json:"seat_type"`
	Price     float64 `json:"price"`
}

package movieinterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	theaterdomain "bookingcinema/pkg/movie/domain"
	"context"
)

type ITheaterRepository interface {
	ListTheaters(ctx context.Context) ([]theaterdomain.Theater, error)
	GetScreens(ctx context.Context, theaterID uint) ([]theaterdomain.Screen, error)
	AddShowtime(ctx context.Context, showtime *bookingdomain.Showtime) error
	GetShowtimes(ctx context.Context, movieID uint, date string) ([]*bookingdomain.Showtime, error)
}

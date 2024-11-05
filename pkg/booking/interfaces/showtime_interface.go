package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type IShowtimeRepository interface {
	GetShowtimeByID(ctx context.Context, id uint) (*bookingdomain.Showtime, error)
	ListShowtimesByMovie(ctx context.Context, movieID uint) ([]*bookingdomain.Showtime, error)
	CreateShowtime(ctx context.Context, showtime *bookingdomain.Showtime) error
}

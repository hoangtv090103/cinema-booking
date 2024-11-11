package theaterinterface

import (
	bookingdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

type IShowtimeRepository interface {
	GetByID(ctx context.Context, id uint) (*bookingdomain.Showtime, error)
	GetByMovie(ctx context.Context, movieID uint) ([]*bookingdomain.Showtime, error)
	Create(ctx context.Context, showtime *bookingdomain.Showtime) error
}

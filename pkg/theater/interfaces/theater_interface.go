package theaterinterface

import (
	moviedomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

// ITheaterRepository TheaterRepository defines methods for theater data access.
type ITheaterRepository interface {
	GetByID(ctx context.Context, id uint) (*moviedomain.Theater, error)
	GetByName(ctx context.Context, name string) ([]*moviedomain.Theater, error)
	GetAll(ctx context.Context) ([]*moviedomain.Theater, error)
	Create(ctx context.Context, theater *moviedomain.Theater) error
	Update(ctx context.Context, theater *moviedomain.Theater) error
	Delete(ctx context.Context, id uint) error
}

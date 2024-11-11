package theaterinfra

import (
	moviedomain "bookingcinema/pkg/theater/domain"
	"context"
)

// TheaterRepository defines methods for theater data access.
type TheaterRepository interface {
	GetByID(ctx context.Context, id uint) (*moviedomain.Theater, error)
	GetAll(ctx context.Context) ([]*moviedomain.Theater, error)
	Create(ctx context.Context, theater *moviedomain.Theater) error
	Update(ctx context.Context, theater *moviedomain.Theater) error
	Delete(ctx context.Context, id uint) error
}

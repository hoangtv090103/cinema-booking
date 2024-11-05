package theaterinterface

import (
	theaterdomain "bookingcinema/pkg/theater/domain"
	"context"
)

type ITheaterRepository interface {
	GetTheaterByID(ctx context.Context, id uint) (*theaterdomain.Theater, error)
	ListTheaters(ctx context.Context) ([]*theaterdomain.Theater, error)
	CreateTheater(ctx context.Context, theater *theaterdomain.Theater) ([]*theaterdomain.Theater, error)
}

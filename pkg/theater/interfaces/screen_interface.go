package theaterinterface

import (
	theaterdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

type IScreenRepository interface {
	GetByID(ctx context.Context, screenID uint) (*theaterdomain.Screen, error)
	GetByTheater(ctx context.Context, theaterID uint) ([]*theaterdomain.Screen, error)
	Create(ctx context.Context, screen *theaterdomain.ScreenCreate) error
}

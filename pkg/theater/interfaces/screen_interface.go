package theaterinterface

import (
	theaterdomain "bookingcinema/pkg/theater/domain"
	"context"
)

type IScreenRepository interface {
	GetScreenByID(ctx context.Context, id uint) (*theaterdomain.Screen, error)
	ListScreensByTheater(ctx context.Context, theaterID uint) ([]*theaterdomain.Screen, error)
	CreateScreen(ctx context.Context, screen *theaterdomain.Screen) error
}

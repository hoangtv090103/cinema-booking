package theaterinterface

import (
	theaterdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

type ISeatRepository interface {
	GetByID(ctx context.Context, id uint) (*theaterdomain.Seat, error)
	GetByScreen(ctx context.Context, screenID uint) ([]*theaterdomain.Seat, error)
	Create(ctx context.Context, seat *theaterdomain.SeatCreate) error
}

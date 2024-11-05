package theaterinterface

import (
	theaterdomain "bookingcinema/pkg/theater/domain"
	"context"
)

type ISeatRepository interface {
	GetSeatByID(ctx context.Context, id uint) (*theaterdomain.Seat, error)
	ListSeatsByScreen(ctx context.Context, screenID uint) ([]*theaterdomain.Seat, error)
	CreateSeat(ctx context.Context, seat *theaterdomain.Seat) error
}

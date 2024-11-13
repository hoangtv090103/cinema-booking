package theaterusecases

import (
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	theaterdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

type ISeatUseCase interface {
	GetSeatByID(ctx context.Context, seatID uint) (*theaterdomain.Seat, error)
	GetSeatsByShowtime(ctx context.Context, showtimeID uint) ([]*theaterdomain.Seat, error)
	GetSeatsByScreen(ctx context.Context, screenID uint) ([]*theaterdomain.Seat, error)
	CreateSeat(ctx context.Context, seat *theaterdomain.SeatCreate) error
}

type SeatUseCase struct {
	repo theaterinterface.ISeatRepository
}

func NewSeatUseCase(repo theaterinterface.ISeatRepository) ISeatUseCase {
	return &SeatUseCase{repo: repo}
}

func (uc *SeatUseCase) GetSeatByID(ctx context.Context, seatID uint) (*theaterdomain.Seat, error) {
	return uc.repo.GetByID(ctx, seatID)
}

func (uc *SeatUseCase) GetSeatsByShowtime(ctx context.Context, showtimeID uint) ([]*theaterdomain.Seat, error) {
	return uc.repo.GetByShowtime(ctx, showtimeID)
}

func (uc *SeatUseCase) GetSeatsByScreen(ctx context.Context, screenID uint) ([]*theaterdomain.Seat, error) {
	return uc.repo.GetByScreen(ctx, screenID)
}

func (uc *SeatUseCase) CreateSeat(ctx context.Context, seat *theaterdomain.SeatCreate) error {
	return uc.repo.Create(ctx, seat)
}

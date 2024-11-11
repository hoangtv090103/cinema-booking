package theaterusecases

import (
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	bookingdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

type IShowtimeUseCase interface {
	GetShowtimeByID(ctx context.Context, showtimeID uint) (*bookingdomain.Showtime, error)
	GetShowtimesByMovie(ctx context.Context, movieID uint) ([]*bookingdomain.Showtime, error)
	CreateShowtime(ctx context.Context, showtime *bookingdomain.Showtime) error
}

type ShowtimeUseCase struct {
	repo theaterinterface.IShowtimeRepository
}

func NewShowtimeUseCase(repo theaterinterface.IShowtimeRepository) IShowtimeUseCase {
	return &ShowtimeUseCase{repo: repo}
}

func (uc *ShowtimeUseCase) GetShowtimeByID(ctx context.Context, showtimeID uint) (*bookingdomain.Showtime, error) {
	return uc.repo.GetByID(ctx, showtimeID)
}

func (uc *ShowtimeUseCase) GetShowtimesByMovie(ctx context.Context, movieID uint) ([]*bookingdomain.Showtime, error) {
	return uc.repo.GetByMovie(ctx, movieID)
}

func (uc *ShowtimeUseCase) CreateShowtime(ctx context.Context, showtime *bookingdomain.Showtime) error {
	return uc.repo.Create(ctx, showtime)
}

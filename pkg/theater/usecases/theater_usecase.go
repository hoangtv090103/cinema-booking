package theaterusecases

import (
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	moviedomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

type ITheaterUseCase interface {
	GetTheaterByID(ctx context.Context, theaterID uint) (*moviedomain.Theater, error)
	GetAllTheaters(ctx context.Context) ([]*moviedomain.Theater, error)
	CreateTheater(ctx context.Context, theater *moviedomain.Theater) error
	UpdateTheater(ctx context.Context, theater *moviedomain.Theater) error
	DeleteTheater(ctx context.Context, theaterID uint) error
}

type TheaterUseCase struct {
	repo theaterinterface.ITheaterRepository
}

func NewTheaterUseCase(repo theaterinterface.ITheaterRepository) ITheaterUseCase {
	return &TheaterUseCase{repo: repo}
}

func (uc *TheaterUseCase) GetTheaterByID(ctx context.Context, theaterID uint) (*moviedomain.Theater, error) {
	return uc.repo.GetByID(ctx, theaterID)
}

func (uc *TheaterUseCase) GetTheaterByName(ctx context.Context, name string) ([]*moviedomain.Theater, error) {
	return uc.repo.GetByName(ctx, name)
}

func (uc *TheaterUseCase) GetAllTheaters(ctx context.Context) ([]*moviedomain.Theater, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *TheaterUseCase) CreateTheater(ctx context.Context, theater *moviedomain.Theater) error {
	return uc.repo.Create(ctx, theater)
}

func (uc *TheaterUseCase) UpdateTheater(ctx context.Context, theater *moviedomain.Theater) error {
	return uc.repo.Update(ctx, theater)
}

func (uc *TheaterUseCase) DeleteTheater(ctx context.Context, theaterID uint) error {
	return uc.repo.Delete(ctx, theaterID)
}

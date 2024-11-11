package theaterusecases

import (
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	moviedomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
)

type IScreenUseCase interface {
	GetScreenByID(ctx context.Context, screenID uint) (*moviedomain.Screen, error)
	GetScreensByTheater(ctx context.Context, theaterID uint) ([]*moviedomain.Screen, error)
	CreateScreen(ctx context.Context, screen *moviedomain.ScreenCreate) error
}

type ScreenUseCase struct {
	repo theaterinterface.IScreenRepository
}

func NewScreenUseCase(repo theaterinterface.IScreenRepository) IScreenUseCase {
	return &ScreenUseCase{repo: repo}
}

func (uc *ScreenUseCase) GetScreenByID(ctx context.Context, screenID uint) (*moviedomain.Screen, error) {
	return uc.repo.GetByID(ctx, screenID)
}

func (uc *ScreenUseCase) GetScreensByTheater(ctx context.Context, theaterID uint) ([]*moviedomain.Screen, error) {
	return uc.repo.GetByTheater(ctx, theaterID)
}

func (uc *ScreenUseCase) CreateScreen(ctx context.Context, screen *moviedomain.ScreenCreate) error {
	return uc.repo.Create(ctx, screen)
}

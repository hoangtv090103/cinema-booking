package movieusecases

import (
	moviedomain "bookingcinema/pkg/movie/domain"
	movieinterface "bookingcinema/pkg/movie/interfaces"
	"context"
)

type IMovieUseCase interface {
	GetMoviesByName(ctx context.Context, name string) ([]*moviedomain.Movie, error)
	GetAllMovies(ctx context.Context) ([]*moviedomain.Movie, error)
	CreateMovie(ctx context.Context, movie *moviedomain.Movie) error
	DeleteMovie(ctx context.Context, movieID uint) error
}

type MovieUseCase struct {
	repo movieinterface.IMovieRepository
}

func NewMovieUseCase(repo movieinterface.IMovieRepository) IMovieUseCase {
	return &MovieUseCase{repo: repo}
}

func (uc *MovieUseCase) GetMoviesByName(ctx context.Context, name string) ([]*moviedomain.Movie, error) {
	return uc.repo.GetByName(ctx, name)
}

func (uc *MovieUseCase) GetAllMovies(ctx context.Context) ([]*moviedomain.Movie, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *MovieUseCase) CreateMovie(ctx context.Context, movie *moviedomain.Movie) error {
	return uc.repo.Create(ctx, movie)
}

func (uc *MovieUseCase) DeleteMovie(ctx context.Context, movieID uint) error {
	return uc.repo.Delete(ctx, movieID)
} 
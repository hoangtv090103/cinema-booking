//go:generate mockgen -source=movie_interface.go -destination=../mocks/theater.go

package interfaces

import (
	moviedomain "bookingcinema/pkg/movie/domain"
	"context"
)

type IMovieRepository interface {
	GetByName(ctx context.Context, name string) ([]*moviedomain.Movie, error)
	GetAll(ctx context.Context) ([]*moviedomain.Movie, error)
	Create(ctx context.Context, movie *moviedomain.Movie) error
	Delete(ctx context.Context, movieID uint) error
}

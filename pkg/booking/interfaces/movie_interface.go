//go:generate mockgen -source=movie_interface.go -destination=../mocks/movie.go

package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type IMovieRepository interface {
	GetMovieByID(ctx context.Context, id uint) (*bookingdomain.Movie, error)
	GetRelativeMoviesByName(ctx context.Context, name string) ([]*bookingdomain.Movie, error)
	ListMovies(ctx context.Context) ([]*bookingdomain.Movie, error)
	CreateMovie(ctx context.Context, movie *bookingdomain.Movie) error
}

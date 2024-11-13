package movieinfra

import (
	moviedomain "bookingcinema/pkg/movie/domain"
	movieinterface "bookingcinema/pkg/movie/interfaces"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) movieinterface.IMovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetAll(ctx context.Context) ([]*moviedomain.Movie, error) {
	var (
		movies []*moviedomain.Movie
	)

	query := `
		SELECT id, title, description, release_date, duration FROM movies WHERE active = true
	`

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("cannot get theater list: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var movie moviedomain.Movie

		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.ReleaseDate,
			&movie.Duration,
		)

		if err != nil {
			return nil, fmt.Errorf("cannot fetch theater data: %v", err)
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (r *MovieRepository) GetByName(ctx context.Context, name string) ([]*moviedomain.Movie, error) {
	var movies []*moviedomain.Movie

	query := `
		SELECT id, title, description, release_date, duration, created_at, updated_at, active
		FROM movies
		WHERE to_tsvector('english', title) @@ to_tsquery('english', $1) AND active = true
	`

	rows, err := r.db.QueryContext(ctx, query, name)
	if err != nil {
		return nil, fmt.Errorf("cannot get relative movies by name: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var movie moviedomain.Movie
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.ReleaseDate,
			&movie.Duration,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.Active,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan theater: %v", err)
		}
		movies = append(movies, &movie)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return movies, nil
}

func (r *MovieRepository) GetByID(ctx context.Context, movieID uint) (*moviedomain.Movie, error) {
	var movie moviedomain.Movie

	query := `
		SELECT id, title, description, release_date, duration, created_at, updated_at, active
		FROM movies
		WHERE id = $1 AND active = true
	`

	err := r.db.QueryRowContext(ctx, query, movieID).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.ReleaseDate,
		&movie.Duration,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.Active,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot get theater by id: %v", err)
	}

	return &movie, nil
}

func (r *MovieRepository) Create(ctx context.Context, movie *moviedomain.Movie) error {
	query := `
		INSERT INTO movies (title, description, release_date, duration)
		VALUES (?, ?, ?, ?)
		RETURNING id
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	err := r.db.QueryRowContext(ctx, query, movie.Title, movie.Description, movie.ReleaseDate).Scan(&movie.ID)
	if err != nil {
		return fmt.Errorf("cannot create theater: %v", err)
	}
	return nil
}

func (r *MovieRepository) Delete(ctx context.Context, movieID uint) error {
	// Soft Delete
	query := `
		UPDATE movies
		SET active = false
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, movieID)
	if err != nil {
		return fmt.Errorf("cannot delete theater with id %d: %v", movieID, err)
	}
	return nil
}

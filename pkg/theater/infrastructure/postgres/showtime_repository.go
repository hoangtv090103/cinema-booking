package theaterinfra

import (
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	bookingdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ShowtimeRepository struct {
	db *sql.DB
}

func NewShowtimeRepository(db *sql.DB) theaterinterface.IShowtimeRepository {
	return &ShowtimeRepository{
		db: db,
	}
}

func (r *ShowtimeRepository) GetByID(ctx context.Context, id uint) (*bookingdomain.Showtime, error) {
	var showtime bookingdomain.Showtime

	query := `
		SELECT id, movie_id, screen_id, start_time, created_at, updated_at, active
		FROM showtimes
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&showtime.ID,
		&showtime.MovieID,
		&showtime.ScreenID,
		&showtime.StartTime,
		&showtime.CreatedAt,
		&showtime.UpdatedAt,
		&showtime.Active,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot get showtime with id %d: %v", id, err)
	}

	return &showtime, nil
}
func (r *ShowtimeRepository) GetByMovie(ctx context.Context, movieID uint) ([]*bookingdomain.Showtime, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, movie_id, screen_id, start_time, created_at, updated_at, active FROM showtimes WHERE movie_id = $1 AND active = true",
		movieID,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot get showtimes for movie %d: %v", movieID, err)
	}
	defer rows.Close()

	var showtimes []*bookingdomain.Showtime
	for rows.Next() {
		var showtime bookingdomain.Showtime
		if err := rows.Scan(&showtime.ID, &showtime.MovieID, &showtime.ScreenID, &showtime.StartTime, &showtime.CreatedAt, &showtime.UpdatedAt, &showtime.Active); err != nil {
			return nil, fmt.Errorf("cannot scan showtime: %v", err)
		}
		showtimes = append(showtimes, &showtime)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return showtimes, nil
}
func (r *ShowtimeRepository) Create(ctx context.Context, showtime *bookingdomain.Showtime) error {
	query := "INSERT INTO showtimes (movie_id, screen_id, start_time, created_at, updated_at, active) VALUES (?, ?, ?, ?, ?, ?, ?)"

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	_, err := r.db.ExecContext(ctx, query, showtime.MovieID, showtime.ScreenID, showtime.StartTime, showtime.CreatedAt, showtime.UpdatedAt, showtime.Active)
	if err != nil {
		return fmt.Errorf("cannot create showtime: %v", err)
	}

	return nil
}

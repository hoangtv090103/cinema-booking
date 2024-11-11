package theaterinfra

import (
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	theaterdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
	"database/sql"
	"fmt"
)

type SeatRepository struct {
	db *sql.DB
}

func NewSeatRepository(db *sql.DB) theaterinterface.ISeatRepository {
	return &SeatRepository{
		db: db,
	}
}

func (r *SeatRepository) GetByID(ctx context.Context, id uint) (*theaterdomain.Seat, error) {
	var seat theaterdomain.Seat

	query := `
		SELECT id, screen_id, row, number, created_at, updated_at, active
		FROM seats
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&seat.ID,
		&seat.ScreenID,
		&seat.Row,
		&seat.Number,
		&seat.CreatedAt,
		&seat.UpdatedAt,
		&seat.Active,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot get seat with id %d: %v", id, err)
	}

	return &seat, nil
}

func (r *SeatRepository) GetByScreen(ctx context.Context, screenID uint) ([]*theaterdomain.Seat, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, screen_id, row, number, created_at, updated_at, active FROM seats WHERE screen_id = $1 AND active = true", screenID)
	if err != nil {
		return nil, fmt.Errorf("cannot get seats for screen %d: %v", screenID, err)
	}
	defer rows.Close()

	var seats []*theaterdomain.Seat
	for rows.Next() {
		var seat theaterdomain.Seat
		if err := rows.Scan(&seat.ID, &seat.ScreenID, &seat.Row, &seat.Number, &seat.CreatedAt, &seat.UpdatedAt, &seat.Active); err != nil {
			return nil, fmt.Errorf("cannot scan seat: %v", err)
		}
		seats = append(seats, &seat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return seats, nil
}

func (r *SeatRepository) Create(ctx context.Context, seat *theaterdomain.SeatCreate) error {
	query := "INSERT INTO seats (screen_id, row, number, created_at, updated_at, active) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.ExecContext(ctx, query, seat.ScreenID, seat.Row, seat.Number)
	if err != nil {
		return fmt.Errorf("cannot create seat: %v", err)
	}
	return nil
}

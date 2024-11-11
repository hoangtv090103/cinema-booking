package theaterinfra

import (
	movieinterface "bookingcinema/pkg/theater/interfaces"
	theaterdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ScreenRepository struct {
	db *sql.DB
}

func NewScreenRepository(db *sql.DB) movieinterface.IScreenRepository {
	return &ScreenRepository{
		db: db,
	}
}

func (r *ScreenRepository) GetByID(ctx context.Context, screenID uint) (*theaterdomain.Screen, error) {
	var screen theaterdomain.Screen

	query := `
		SELECT id, theater_id, name, capacity 
		FROM screens 
		WHERE id = ?
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	err := r.db.QueryRowContext(ctx, query, screenID).Scan(
		&screen.ID,
		&screen.TheaterID,
		&screen.ScreenName,
		&screen.Capacity,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot get screen by id %d: %v", screenID, err)
	}

	return &screen, nil
}

func (r *ScreenRepository) GetByTheater(ctx context.Context, theaterID uint) ([]*theaterdomain.Screen, error) {
	var (
		screens []*theaterdomain.Screen
	)

	query := `
		SELECT id, theater_id, name, capacity 
		FROM screens 
		WHERE theater_id = ?
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	rows, err := r.db.QueryContext(ctx, query, theaterID)
	if err != nil {
		return nil, fmt.Errorf("cannot get screens of theater with id %d: %v", theaterID, err)
	}

	defer rows.Close()

	for rows.Next() {
		var screen theaterdomain.Screen
		err = rows.Scan(
			&screen.ID,
			&screen.TheaterID,
			&screen.ScreenName,
			&screen.Capacity,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot get screens of theater with id %d: %v", theaterID, err)
		}

		screens = append(screens, &screen)
	}

	return screens, nil
}
func (r *ScreenRepository) Create(ctx context.Context, screen *theaterdomain.ScreenCreate) error {
	query := `
		INSERT INTO screens (theater_id, name, capacity) 
		VALUES (?, ?, ?)
	`

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	_, err := r.db.ExecContext(ctx, query, screen.TheaterID, screen.ScreenName, screen.Capacity)

	if err != nil {
		return fmt.Errorf("cannot insert screen %s: %v", screen.ScreenName, err)
	}

	return nil
}

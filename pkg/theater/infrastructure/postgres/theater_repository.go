package theaterinfra

import (
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	theaterdomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
	"database/sql"
	"fmt"
)

type TheaterRepository struct {
	db *sql.DB
}

func NewTheaterRepository(db *sql.DB) theaterinterface.ITheaterRepository {
	return &TheaterRepository{db: db}
}

func (r *TheaterRepository) GetByID(ctx context.Context, id uint) (*theaterdomain.Theater, error) {
	var theater theaterdomain.Theater
	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, location, created_at, updated_at, active FROM theaters WHERE id = $1 AND active = true",
		id,
	).Scan(
		&theater.ID,
		&theater.Name,
		&theater.Location,
		&theater.CreatedAt,
		&theater.UpdatedAt,
		&theater.Active,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("theater not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting theater: %v", err)
	}

	return &theater, nil
}

func (r *TheaterRepository) GetByName(ctx context.Context, name string) ([]*theaterdomain.Theater, error) {
	var theaters []*theaterdomain.Theater
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, location, created_at, updated_at, active FROM theaters WHERE name = $1 AND active = true",
		name,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting theaters: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var theater theaterdomain.Theater
		if err := rows.Scan(
			&theater.ID,
			&theater.Name,
			&theater.Location,
			&theater.CreatedAt,
			&theater.UpdatedAt,
			&theater.Active,
		); err != nil {
			return nil, fmt.Errorf("error scanning theater: %v", err)
		}
	}

	return theaters, nil
}

func (r *TheaterRepository) GetAll(ctx context.Context) ([]*theaterdomain.Theater, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, location, created_at, updated_at, active FROM theaters WHERE active = true",
	)
	if err != nil {
		return nil, fmt.Errorf("error getting theaters: %v", err)
	}
	defer rows.Close()

	var theaters []*theaterdomain.Theater
	for rows.Next() {
		var theater theaterdomain.Theater
		if err := rows.Scan(
			&theater.ID,
			&theater.Name,
			&theater.Location,
			&theater.CreatedAt,
			&theater.UpdatedAt,
			&theater.Active,
		); err != nil {
			return nil, fmt.Errorf("error scanning theater: %v", err)
		}
		theaters = append(theaters, &theater)
	}

	return theaters, nil
}

func (r *TheaterRepository) Create(ctx context.Context, theater *theaterdomain.Theater) error {
	return r.db.QueryRowContext(
		ctx,
		"INSERT INTO theaters (name, location) VALUES ($1, $2) RETURNING id",
		theater.Name,
		theater.Location,
	).Scan(&theater.ID)
}

func (r *TheaterRepository) Update(ctx context.Context, theater *theaterdomain.Theater) error {
	result, err := r.db.ExecContext(
		ctx,
		"UPDATE theaters SET name = $1, location = $2, updated_at = NOW() WHERE id = $3 AND active = true",
		theater.Name,
		theater.Location,
		theater.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating theater: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting affected rows: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("theater not found")
	}

	return nil
}

func (r *TheaterRepository) Delete(ctx context.Context, id uint) error {
	result, err := r.db.ExecContext(
		ctx,
		"UPDATE theaters SET active = false, updated_at = NOW() WHERE id = $1 AND active = true",
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting theater: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting affected rows: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("theater not found")
	}

	return nil
}

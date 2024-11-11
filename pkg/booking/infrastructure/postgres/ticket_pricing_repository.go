package bookinginfra

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	bookinginterface "bookingcinema/pkg/booking/interfaces"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TicketPricingRepository struct {
	db *sql.DB
}

func NewTicketPricingRepository(db *sql.DB) bookinginterface.ITicketPricingRepository {
	return &TicketPricingRepository{
		db: db,
	}
}

func (r *TicketPricingRepository) GetAll(ctx context.Context) ([]*bookingdomain.TicketPricing, error) {
	var ticketPricings []*bookingdomain.TicketPricing

	query := `SELECT id, seat_type, day_of_week, price, created_at, updated_at, active FROM ticket_pricing`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ticketPricing bookingdomain.TicketPricing

		err = rows.Scan(
			&ticketPricing.ID,
			&ticketPricing.SeatType,
			&ticketPricing.DayOfWeek,
			&ticketPricing.Price,
			&ticketPricing.CreatedAt,
			&ticketPricing.UpdatedAt,
			&ticketPricing.Active,
		)

		if err != nil {
			return nil, fmt.Errorf("error when getting ticket price: %v", err)
		}

		ticketPricings = append(ticketPricings, &ticketPricing)
	}
	return ticketPricings, nil
}

func (r *TicketPricingRepository) GetByID(ctx context.Context, id uint) (*bookingdomain.TicketPricing, error) {
	var ticketPricing bookingdomain.TicketPricing
	query := `
		SELECT id, seat_type, day_of_week, price, created_at, updated_at, active 
		FROM ticket_pricing  
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&ticketPricing.ID,
		&ticketPricing.SeatType,
		&ticketPricing.DayOfWeek,
		&ticketPricing.Price,
		&ticketPricing.CreatedAt,
		&ticketPricing.UpdatedAt,
		&ticketPricing.Active,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot get ticket price with id %d: %v", id, err)
	}

	return &ticketPricing, nil
}

func (r *TicketPricingRepository) GetBySeatType(ctx context.Context, seatType string) ([]*bookingdomain.TicketPricing, error) {
	var ticketPricings []*bookingdomain.TicketPricing
	query := `
		SELECT id, seat_type, day_of_week, price, created_at, updated_at, active 
		FROM ticket_pricing  
		WHERE seat_type = $1
	`

	rows, err := r.db.QueryContext(ctx, query, seatType)
	if err != nil {
		return nil, fmt.Errorf("error when getting ticket price by seat type: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ticketPricing bookingdomain.TicketPricing

		err = rows.Scan(
			&ticketPricing.ID,
			&ticketPricing.SeatType,
			&ticketPricing.DayOfWeek,
			&ticketPricing.Price,
			&ticketPricing.CreatedAt,
			&ticketPricing.UpdatedAt,
			&ticketPricing.Active,
		)

		if err != nil {
			return nil, fmt.Errorf("error when getting ticket price: %v", err)
		}

		ticketPricings = append(ticketPricings, &ticketPricing)
	}
	return ticketPricings, nil
}

func (r *TicketPricingRepository) GetByDayAndType(ctx context.Context, dayOfWeek uint, seatType string) (*bookingdomain.TicketPricing, error) {
	var ticketPricing bookingdomain.TicketPricing
	query := `
		SELECT id, seat_type, day_of_week, price, created_at, updated_at, active 
		FROM ticket_pricing  
		WHERE day_of_week = ? AND seat_type = ?
	`

	err := r.db.QueryRowContext(ctx, query, dayOfWeek, seatType).Scan(
		&ticketPricing.ID,
		&ticketPricing.SeatType,
		&ticketPricing.DayOfWeek,
		&ticketPricing.Price,
		&ticketPricing.CreatedAt,
		&ticketPricing.UpdatedAt,
		&ticketPricing.Active,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot get ticket price with day of week %d and seat type %s: %v", dayOfWeek, seatType, err)
	}

	return &ticketPricing, nil
}

func (r *TicketPricingRepository) Create(ctx context.Context, pricing *bookingdomain.TicketPricing) error {
	query := `
		INSERT INTO ticket_pricing (seat_type, day_of_week, price, created_at, updated_at, active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error when preparing insert query: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		&pricing.SeatType,
		&pricing.DayOfWeek,
		&pricing.Price,
		&pricing.CreatedAt,
		&pricing.UpdatedAt,
		&pricing.Active,
	).Scan(&pricing.ID)

	if err != nil {
		return fmt.Errorf("error when inserting ticket price: %v", err)
	}

	return nil
}

func (r *TicketPricingRepository) Update(ctx context.Context, pricing *bookingdomain.TicketPricing) error {
	query := `
		UPDATE ticket_pricing
		SET seat_type = ?, day_of_week = ?, price = ?, updated_at = ?, active = ?
		WHERE id = ?
	`

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error when preparing update query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		&pricing.SeatType,
		&pricing.DayOfWeek,
		&pricing.Price,
		&pricing.UpdatedAt,
		&pricing.Active,
		&pricing.ID,
	)

	if err != nil {
		return fmt.Errorf("error when updating ticket price: %v", err)
	}

	return nil
}

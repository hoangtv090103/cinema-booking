package bookinginfra

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	moviedomain "bookingcinema/pkg/theater/theaterdomain"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) CreateBooking(ctx context.Context, booking *bookingdomain.BookingCreate) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	bookingID, err := r.insertBooking(ctx, booking, tx)
	if err != nil {
		return err
	}

	seats, err := r.GetSeats(ctx, booking.SeatIDs, tx)
	if err != nil {
		return err
	}

	bookingSeats, err := r.CreateBookingSeats(ctx, seats, bookingID, tx)
	if err != nil {
		return err
	}

	totalPrice := r.calculateTotalPrice(bookingSeats)
	err = r.updateBookingTotalPrice(ctx, bookingID, totalPrice, tx)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) insertBooking(ctx context.Context, booking *bookingdomain.BookingCreate, tx *sql.Tx) (uint, error) {
	query := `
		INSERT INTO bookings (user_id, showtime_id, status) VALUES (?, ?, ?) RETURNING id
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	var bookingID uint
	err := tx.QueryRowContext(ctx, query, booking.UserID, booking.ShowtimeID, booking.Status).Scan(&bookingID)
	if err != nil {
		return 0, fmt.Errorf("cannot insert booking: %v", err)
	}
	return bookingID, nil
}

func (r *BookingRepository) GetSeats(ctx context.Context, seatIDs []uint, tx *sql.Tx) ([]*moviedomain.Seat, error) {
	var seats []*moviedomain.Seat
	query := `
		SELECT id, screen_id, row, number, seat_type, available FROM seats
		WHERE id = ANY($1)
	`
	rows, err := tx.QueryContext(ctx, query, pq.Array(seatIDs))
	if err != nil {
		return nil, fmt.Errorf("cannot get seats: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		seat := &moviedomain.Seat{}
		err = rows.Scan(&seat.ID, &seat.ScreenID, &seat.Row, &seat.Number, &seat.SeatType, &seat.Available)
		if err != nil {
			return nil, fmt.Errorf("cannot scan seat: %v", err)
		}
		seats = append(seats, seat)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return seats, nil
}

func (r *BookingRepository) CreateBookingSeats(ctx context.Context, seats []*moviedomain.Seat, bookingID uint, tx *sql.Tx) ([]*bookingdomain.BookingSeat, error) {
	var bookingSeats []*bookingdomain.BookingSeat
	query := `
		INSERT INTO booking_seats (booking_id, seat_id, price) VALUES (?, ?, ?) RETURNING id, booking_id, seat_id, price
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	for _, seat := range seats {
		price, err := r.getSeatPrice(ctx, seat.SeatType, tx)
		if err != nil {
			return nil, fmt.Errorf("cannot get seat price: %v", err)
		}

		var bookingSeat bookingdomain.BookingSeat
		err = tx.QueryRowContext(ctx, query, bookingID, seat.ID, price).Scan(&bookingSeat.ID, &bookingSeat.BookingID, &bookingSeat.SeatID, &bookingSeat.Price)
		if err != nil {
			return nil, fmt.Errorf("cannot insert booking seat: %v", err)
		}

		bookingSeats = append(bookingSeats, &bookingSeat)
	}

	return bookingSeats, nil
}

func (r *BookingRepository) getSeatPrice(ctx context.Context, seatType string, tx *sql.Tx) (float64, error) {
	query := `
		SELECT price FROM ticket_pricing
		WHERE seat_type = ? AND day_of_week = EXTRACT(DOW FROM NOW())
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	var price float64
	err := tx.QueryRowContext(ctx, query, seatType).Scan(&price)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("cannot get seat price: %v", err)
	}
	return price, nil
}

func (r *BookingRepository) calculateTotalPrice(bookingSeats []*bookingdomain.BookingSeat) float64 {
	totalPrice := 0.0
	for _, bookingSeat := range bookingSeats {
		totalPrice += bookingSeat.Price
	}
	return totalPrice
}

func (r *BookingRepository) updateBookingTotalPrice(ctx context.Context, bookingID uint, totalPrice float64, tx *sql.Tx) error {
	query := `
		UPDATE bookings
		SET total_price = $1
		WHERE id = $2
	`
	_, err := tx.ExecContext(ctx, query, totalPrice, bookingID)
	if err != nil {
		return fmt.Errorf("cannot update booking total price: %v", err)
	}
	return nil
}

func (r *BookingRepository) GetUserBookings(ctx context.Context, userID uint) ([]*bookingdomain.Booking, error) {
	var (
		bookings []*bookingdomain.Booking
	)

	query := `
		SELECT b.id, b.user_id, b.showtime_id, b.booking_time, b.total_price, b.created_at, b.updated_at, b.active
		FROM bookings b
		WHERE b.user_id = $1
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var booking bookingdomain.Booking
		err = rows.Scan(
			&booking.ID,
			&booking.UserID,
			&booking.ShowtimeID,
			&booking.BookingTime,
			&booking.TotalPrice,
			&booking.CreatedAt,
			&booking.UpdatedAt,
			&booking.Active,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepository) GetByID(bookingID uint) (*bookingdomain.Booking, error) {
	query := `
		SELECT b.id, b.user_id, b.showtime_id, b.booking_time, b.total_price, b.created_at, b.updated_at, b.active
		FROM bookings b
		LEFT JOIN booking_seats bs ON b.id = bs.booking_id
		WHERE b.id = $1

	`

	var booking bookingdomain.Booking
	err := r.db.QueryRow(query, bookingID).Scan(
		&booking.ID,
		&booking.UserID,
		&booking.ShowtimeID,
		&booking.BookingTime,
		&booking.TotalPrice,
		&booking.CreatedAt,
		&booking.UpdatedAt,
		&booking.Active,
	)

	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *BookingRepository) ConfirmBooking(ctx context.Context, bookingID uint) error {
	query := `
		UPDATE bookings
		SET status = 'confirmed'
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, bookingID)
	if err != nil {
		return fmt.Errorf("cannot confirm booking: %v", err)
	}
	return nil
}

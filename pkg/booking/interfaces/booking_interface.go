package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type IBookingRepository interface {
	GetBookingByID(ctx context.Context, id uint) (*bookingdomain.Booking, error)
	ListBookingsByUser(ctx context.Context, userID uint) ([]*bookingdomain.Booking, error)
	CreateBooking(ctx context.Context, booking *bookingdomain.Booking) error
}

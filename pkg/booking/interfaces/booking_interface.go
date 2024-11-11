package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type IBookingRepository interface {
	CreateBooking(ctx context.Context, booking *bookingdomain.BookingCreate) error
	GetUserBookings(ctx context.Context, userID uint) ([]*bookingdomain.Booking, error)
	ConfirmBooking(ctx context.Context, bookingID uint) error
}

package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type IBookingSeatRepository interface {
	GetBookingSeatByID(ctx context.Context, id uint) (*bookingdomain.BookingSeat, error)
	ListBookingSeatsByBooking(ctx context.Context, bookingID uint) ([]*bookingdomain.BookingSeat, error)
	CreateBookingSeat(ctx context.Context, bookingSeat *bookingdomain.BookingSeat) error
}

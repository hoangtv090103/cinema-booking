package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type IBookingSeatRepository interface {
	GetByBookingID(ctx context.Context, bookingID uint) (*bookingdomain.BookingSeat, error)
	GetBySeatID(ctx context.Context, bookingID uint) (*bookingdomain.BookingSeat, error)
	CreateBookingSeat(ctx context.Context, bookingSeat *bookingdomain.BookingSeat) error
}

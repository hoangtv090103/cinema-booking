package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type IPaymentRepository interface {
	GetPaymentByID(ctx context.Context, id uint) (*bookingdomain.Payment, error)
	ListPaymentsByBooking(ctx context.Context, bookingID uint) ([]*bookingdomain.Payment, error)
	CreatePayment(ctx context.Context, payment *bookingdomain.Payment) error
	UpdatePaymentStatus(ctx context.Context, id uint, status string) error
}

package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type ITicketPricingRepository interface {
	GetAll(ctx context.Context) ([]*bookingdomain.TicketPricing, error)
	GetByID(ctx context.Context, id uint) (*bookingdomain.TicketPricing, error)
	GetBySeatType(ctx context.Context, seatType string) ([]*bookingdomain.TicketPricing, error)
	GetByDayAndType(ctx context.Context, dayOfWeek uint, seatType string) (*bookingdomain.TicketPricing, error)
	Create(ctx context.Context, pricing *bookingdomain.TicketPricing) error
	Update(ctx context.Context, pricing *bookingdomain.TicketPricing) error
}

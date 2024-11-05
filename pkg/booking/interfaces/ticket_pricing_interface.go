package bookinginterface

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	"context"
)

type ITicketPricingRepository interface {
	GetTicketPricingByID(ctx context.Context, id uint) (*bookingdomain.TicketPricing, error)
	ListTicketPricingsBySeatType(ctx context.Context, seatType string) ([]*bookingdomain.TicketPricing, error)
	GetTicketPricingByDayAndType(ctx context.Context, dayOfWeek uint, seatType string) (*bookingdomain.TicketPricing, error)
	CreateTicketPricing(ctx context.Context, pricing *bookingdomain.TicketPricing) error
	UpdateTicketPricing(ctx context.Context, pricing *bookingdomain.TicketPricing) error
}
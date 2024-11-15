package bookinginterface

import (
	"context"
	"database/sql"
)

type IPricingRepository interface {
	CalculateBookingPrice(ctx context.Context, showtimeID uint, seatIDs []uint, tx *sql.Tx) (float64, error)
	GetBasePriceForSeatType(ctx context.Context, seatType string) (float64, error)
	ApplyDiscounts(ctx context.Context, basePrice float64, discountCodes []string) (float64, error)
}
	
package bookingusecases

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	bookinginterface "bookingcinema/pkg/booking/interfaces"
	theaterinterface "bookingcinema/pkg/theater/interfaces"
	"context"
	"fmt"
	"time"
)

type IBookingUseCase interface {
	CreateBooking(ctx context.Context, userID uint, showtimeID uint, seatIDs []uint) (*bookingdomain.Booking, error)
	GetUserBooking(ctx context.Context, userID uint) ([]*bookingdomain.Booking, error)
}
type BookingUseCase struct {
	bookingRepo  bookinginterface.IBookingRepository
	pricingRepo  bookinginterface.ITicketPricingRepository
	showtimeRepo theaterinterface.IShowtimeRepository
	seatRepo     theaterinterface.ISeatRepository
}

func NewBookingUseCase(
	bookingRepo bookinginterface.IBookingRepository,
	pricingRepo bookinginterface.ITicketPricingRepository,
	showtimeRepo theaterinterface.IShowtimeRepository,
	seatRepo theaterinterface.ISeatRepository,
) IBookingUseCase {
	return &BookingUseCase{
		bookingRepo:  bookingRepo,
		pricingRepo:  pricingRepo,
		showtimeRepo: showtimeRepo,
		seatRepo:     seatRepo,
	}
}

func (u *BookingUseCase) CreateBooking(ctx context.Context, userID uint, showtimeID uint, seatIDs []uint) (*bookingdomain.Booking, error) {
	// 1. Verify showtime exists and is valid
	showtime, err := u.showtimeRepo.GetByID(ctx, showtimeID)
	if err != nil {
		return nil, fmt.Errorf("cannot get showtime not found: %v", err)
	}

	if showtime == nil {
		return nil, fmt.Errorf("showtime not found")
	}

	// 2. Create booking
	bookingCreate := &bookingdomain.BookingCreate{
		UserID:      userID,
		ShowtimeID:  showtimeID,
		SeatIDs:     seatIDs,
		Status:      "pending",
		BookingTime: time.Now(),
	}

	err = u.bookingRepo.CreateBooking(ctx, bookingCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking: %v", err)
	}

	// 4. Get user's bookings
	bookings, err := u.bookingRepo.GetUserBookings(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user bookings: %v", err)
	}

	// Return the latest booking
	if len(bookings) > 0 {
		return bookings[len(bookings)-1], nil
	}

	return nil, fmt.Errorf("booking created but not found")
}

func (u *BookingUseCase) GetUserBooking(ctx context.Context, userID uint) ([]*bookingdomain.Booking, error) {
	return u.bookingRepo.GetUserBookings(ctx, userID)
}

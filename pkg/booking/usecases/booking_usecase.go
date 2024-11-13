package bookingusecases

import (
	bookingdomain "bookingcinema/pkg/booking/domain"
	bookingkafka "bookingcinema/pkg/booking/infrastructure/kafka"
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
	producer     *bookingkafka.Producer
}

func NewBookingUseCase(
	bookingRepo bookinginterface.IBookingRepository,
	pricingRepo bookinginterface.ITicketPricingRepository,
	showtimeRepo theaterinterface.IShowtimeRepository,
	seatRepo theaterinterface.ISeatRepository,
	producer *bookingkafka.Producer,
) IBookingUseCase {
	return &BookingUseCase{
		bookingRepo:  bookingRepo,
		pricingRepo:  pricingRepo,
		showtimeRepo: showtimeRepo,
		seatRepo:     seatRepo,
		producer:     producer,
	}
}

func (u *BookingUseCase) CreateBooking(ctx context.Context, userID uint, showtimeID uint, seatIDs []uint) (*bookingdomain.Booking, error) {
	var booking *bookingdomain.Booking
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

	bookings, err := u.bookingRepo.GetUserBookings(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user bookings: %v", err)
	}

	// Return the latest booking
	if len(bookings) > 0 {
		booking = bookings[len(bookings)-1]
	}

	// Publish booking event
	err = u.producer.PublishBookingEvent(booking.ID, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to publish booking event: %v", err)
	}

	return nil, fmt.Errorf("booking created but not found")
}

func (u *BookingUseCase) GetUserBooking(ctx context.Context, userID uint) ([]*bookingdomain.Booking, error) {
	return u.bookingRepo.GetUserBookings(ctx, userID)
}

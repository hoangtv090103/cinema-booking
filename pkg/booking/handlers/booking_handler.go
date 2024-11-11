package bookinghandler

import (
	"bookingcinema/pkg/auth/authdomain"
	"bookingcinema/pkg/booking/domain"
	bookingusecases "bookingcinema/pkg/booking/usecases"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	useCase bookingusecases.IBookingUseCase
}

func NewBookingHandler(useCase bookingusecases.IBookingUseCase) *BookingHandler {
	return &BookingHandler{useCase: useCase}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var bookingCreate bookingdomain.BookingCreate
	if err := c.BodyParser(&bookingCreate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	booking, err := h.useCase.CreateBooking(c.Context(), bookingCreate.UserID, bookingCreate.ShowtimeID, bookingCreate.SeatIDs)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(booking)
}

func (h *BookingHandler) GetUserBookings(c *fiber.Ctx) error {
	userID := c.Locals("user").(*authdomain.User).ID

	bookings, err := h.useCase.GetUserBooking(c.Context(), uint(userID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(bookings)
}

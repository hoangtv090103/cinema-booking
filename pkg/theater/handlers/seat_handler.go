package theaterhandler

import (
	"bookingcinema/pkg/theater/theaterdomain"
	theaterusecases "bookingcinema/pkg/theater/usecases"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SeatHandler struct {
	useCase theaterusecases.ISeatUseCase
}

func NewSeatHandler(useCase theaterusecases.ISeatUseCase) *SeatHandler {
	return &SeatHandler{useCase: useCase}
}

func (h *SeatHandler) GetSeatByID(c *fiber.Ctx) error {
	seatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid seat ID"})
	}
	seat, err := h.useCase.GetSeatByID(c.Context(), uint(seatID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(seat)
}

func (h *SeatHandler) GetSeatsByShowtime(c *fiber.Ctx) error {
	showtimeID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid showtime ID"})
	}

	seats, err := h.useCase.GetSeatsByShowtime(c.Context(), uint(showtimeID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(seats)
}
func (h *SeatHandler) GetSeatsByScreen(c *fiber.Ctx) error {
	screenID, err := strconv.Atoi(c.Params("screen_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid screen ID"})
	}

	seats, err := h.useCase.GetSeatsByScreen(c.Context(), uint(screenID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(seats)
}

func (h *SeatHandler) CreateSeat(c *fiber.Ctx) error {
	var seat theaterdomain.SeatCreate
	if err := c.BodyParser(&seat); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.useCase.CreateSeat(c.Context(), &seat); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(seat)
}

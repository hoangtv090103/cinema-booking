package theaterhandler

import (
	"net/http"
	"strconv"

	"bookingcinema/pkg/theater/theaterdomain"
	theaterusecases "bookingcinema/pkg/theater/usecases"
	"github.com/gofiber/fiber/v2"
)

type ShowtimeHandler struct {
	useCase theaterusecases.IShowtimeUseCase
}

func NewShowtimeHandler(useCase theaterusecases.IShowtimeUseCase) *ShowtimeHandler {
	return &ShowtimeHandler{useCase: useCase}
}

func (h *ShowtimeHandler) GetShowtimeByID(c *fiber.Ctx) error {
	showtimeID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid showtime ID"})
	}
	showtime, err := h.useCase.GetShowtimeByID(c.Context(), uint(showtimeID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(showtime)
}

func (h *ShowtimeHandler) GetShowtimesByMovie(c *fiber.Ctx) error {
	movieID, err := strconv.Atoi(c.Params("movie_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid movie ID"})
	}
	showtimes, err := h.useCase.GetShowtimesByMovie(c.Context(), uint(movieID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(showtimes)
}

func (h *ShowtimeHandler) CreateShowtime(c *fiber.Ctx) error {
	var showtime theaterdomain.Showtime
	if err := c.BodyParser(&showtime); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.useCase.CreateShowtime(c.Context(), &showtime); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(showtime)
}

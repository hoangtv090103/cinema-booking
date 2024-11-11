package theaterhandler

import (
	"bookingcinema/pkg/theater/theaterdomain"
	theaterusecases "bookingcinema/pkg/theater/usecases"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ScreenHandler struct {
	useCase theaterusecases.IScreenUseCase
}

func NewScreenHandler(useCase theaterusecases.IScreenUseCase) *ScreenHandler {
	return &ScreenHandler{useCase: useCase}
}

func (h *ScreenHandler) GetScreenByID(c *fiber.Ctx) error {
	screenID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid screen ID"})
	}
	screen, err := h.useCase.GetScreenByID(c.Context(), uint(screenID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(screen)
}

func (h *ScreenHandler) GetScreensByTheater(c *fiber.Ctx) error {
	theaterID, err := strconv.Atoi(c.Params("theater_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid theater ID"})
	}
	screens, err := h.useCase.GetScreensByTheater(c.Context(), uint(theaterID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(screens)
}

func (h *ScreenHandler) CreateScreen(c *fiber.Ctx) error {
	var screen theaterdomain.ScreenCreate
	if err := c.BodyParser(&screen); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.useCase.CreateScreen(c.Context(), &screen); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(screen)
}

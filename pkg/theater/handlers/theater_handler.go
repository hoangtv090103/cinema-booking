package theaterhandler

import (
	"net/http"
	"strconv"

	"bookingcinema/pkg/theater/theaterdomain"
	theaterusecases "bookingcinema/pkg/theater/usecases"
	"github.com/gofiber/fiber/v2"
)

type TheaterHandler struct {
	useCase theaterusecases.ITheaterUseCase
}

func NewTheaterHandler(useCase theaterusecases.ITheaterUseCase) *TheaterHandler {
	return &TheaterHandler{useCase: useCase}
}

func (h *TheaterHandler) GetTheaterByID(c *fiber.Ctx) error {
	theaterID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid theater ID"})
	}
	theater, err := h.useCase.GetTheaterByID(c.Context(), uint(theaterID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(theater)
}

func (h *TheaterHandler) GetAllTheaters(c *fiber.Ctx) error {
	theaters, err := h.useCase.GetAllTheaters(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(theaters)
}

func (h *TheaterHandler) CreateTheater(c *fiber.Ctx) error {
	var theater theaterdomain.Theater
	if err := c.BodyParser(&theater); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.useCase.CreateTheater(c.Context(), &theater); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(theater)
}

func (h *TheaterHandler) UpdateTheater(c *fiber.Ctx) error {
	var theater theaterdomain.Theater
	if err := c.BodyParser(&theater); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	theaterID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid theater ID"})
	}
	theater.ID = uint(theaterID)
	if err := h.useCase.UpdateTheater(c.Context(), &theater); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(theater)
}

func (h *TheaterHandler) DeleteTheater(c *fiber.Ctx) error {
	theaterID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid theater ID"})
	}
	if err := h.useCase.DeleteTheater(c.Context(), uint(theaterID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}

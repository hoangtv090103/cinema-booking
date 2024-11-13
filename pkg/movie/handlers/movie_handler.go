package moviehandler

import (
	"net/http"
	"strconv"

	moviedomain "bookingcinema/pkg/movie/domain"
	movieusecases "bookingcinema/pkg/movie/usecases"
	"github.com/gofiber/fiber/v2"
)

type MovieHandler struct {
	useCase movieusecases.IMovieUseCase
}

func NewMovieHandler(useCase movieusecases.IMovieUseCase) *MovieHandler {
	return &MovieHandler{useCase: useCase}
}

func (h *MovieHandler) GetAllMovies(c *fiber.Ctx) error {
	movies, err := h.useCase.GetAllMovies(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch movies",
		})
	}
	return c.JSON(movies)
}

func (h *MovieHandler) GetMovieByID(c *fiber.Ctx) error {
	movieID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid movie ID",
		})
	}

	movie, err := h.useCase.GetMovieByID(c.Context(), uint(movieID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Movie not found",
		})
	}
	return c.JSON(movie)
}

func (h *MovieHandler) SearchMovies(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Search term is required",
		})
	}

	movies, err := h.useCase.GetMoviesByName(c.Context(), name)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search movies",
		})
	}
	return c.JSON(movies)
}

func (h *MovieHandler) CreateMovie(c *fiber.Ctx) error {
	movie := new(moviedomain.Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.useCase.CreateMovie(c.Context(), movie); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create movie",
		})
	}

	return c.Status(http.StatusCreated).JSON(movie)
}

func (h *MovieHandler) DeleteMovie(c *fiber.Ctx) error {
	movieID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid movie ID",
		})
	}

	if err := h.useCase.DeleteMovie(c.Context(), uint(movieID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete movie",
		})
	}

	return c.SendStatus(http.StatusNoContent)
} 
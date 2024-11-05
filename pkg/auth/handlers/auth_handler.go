package authhandler

import (
	"bookingcinema/pkg/auth/authdomain"
	authusecase "bookingcinema/pkg/auth/usecases"
	authutils "bookingcinema/pkg/auth/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase authusecase.IAuthenticationUseCase
}

func NewAuthHandler(authUseCase authusecase.IAuthenticationUseCase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUseCase,
	}
}

func (h *AuthHandler) RegisterHandler(c *fiber.Ctx) error {
	var user authdomain.UserCreate

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	err := h.authUsecase.Register(c.Context(), &user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

// LoginHandler logs in a user and returns a JWT token
func (h *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	creds := authdomain.UserLogin{}

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	user, err := h.authUsecase.Login(c.Context(), creds.Email, creds.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := authutils.GenerateJWT(user.ID, user.RoleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// UserHandler returns the details of the authenticated user
func (h *AuthHandler) UserHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*authdomain.User)
	return c.JSON(user)
}

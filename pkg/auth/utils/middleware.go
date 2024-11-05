package authutils

import (
	authinterface "bookingcinema/pkg/auth/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func AuthMiddleware(userRepo authinterface.IUserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))
		roleID := uint(claims["role_id"].(float64))

		user, err := userRepo.GetByID(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		// Store the user in context for the next handler
		c.Locals("user", user)
		c.Locals("role_id", roleID)
		return c.Next()
	}
}

func AdminOnlyMiddleware(roleRepo authinterface.IRoleRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID := uint(c.Locals("role_id").(int64))
		role, err := roleRepo.GetByName(c.Context(), "admin")
		if err != nil {
			return nil
		}
		if roleID != role.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Admin access required"})
		}
		return c.Next()
	}
}

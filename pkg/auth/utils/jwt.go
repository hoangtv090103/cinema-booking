package authutils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRETE_KEY"))

func GenerateJWT(userID uint, roleID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["role_id"] = roleID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token valid for 24 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT parses and validates a JWT token
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

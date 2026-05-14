package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pholguinc/api-go-matrices/internal/dtos"
)

func AuthMiddleware(c fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.NewErrorResponse(fiber.StatusUnauthorized, "Token no proporcionado"))
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.NewErrorResponse(fiber.StatusUnauthorized, "Formato de token inválido"))
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.NewErrorResponse(fiber.StatusUnauthorized, "Token inválido o expirado"))
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", claims["user_id"])

	return c.Next()
}

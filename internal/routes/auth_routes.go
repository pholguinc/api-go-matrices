package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pholguinc/api-go-matrices/internal/controllers"
)

func SetupAuthRoutes(router fiber.Router, ctrl *controllers.AuthController) {
	authGroup := router.Group("/auth")
	authGroup.Post("/register", ctrl.Register)
	authGroup.Post("/login", ctrl.Login)
}

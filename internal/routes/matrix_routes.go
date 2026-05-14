package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pholguinc/api-go-matrices/internal/controllers"
	"github.com/pholguinc/api-go-matrices/internal/middlewares"
)

func SetupMatrixRoutes(router fiber.Router, ctrl *controllers.MatrixController) {
	// Protegemos el grupo de matrices con el AuthMiddleware
	matrixGroup := router.Group("/matrix", middlewares.AuthMiddleware)

	matrixGroup.Post("/factorize", ctrl.Factorize)

	// Las rutas de documentación siguen siendo públicas
	router.Get("/docs/swagger.json", func(c fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	router.Get("/docs", func(c fiber.Ctx) error {
		return c.SendFile("./docs/index.html")
	})
}

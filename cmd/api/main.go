package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/pholguinc/api-go-matrices/internal/controllers"
	"github.com/pholguinc/api-go-matrices/internal/middlewares"
	"github.com/pholguinc/api-go-matrices/internal/routes"
	"github.com/pholguinc/api-go-matrices/internal/services"
)

// @title Matrix Factorization API
// @version 1.0
// @description API para realizar factorización QR de matrices rectangulares.
// @host localhost:3001
// @BasePath /
func main() {
	app := fiber.New()

	// Middlewares
	app.Use(middlewares.Logger)

	// Environments
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// Dependency Injection
	service := services.NewMatrixService()
	controller := controllers.NewMatrixController(service)

	// Routes
	routes.SetupMatrixRoutes(app, controller)

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

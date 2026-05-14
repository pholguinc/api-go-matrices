package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pholguinc/api-go-matrices/internal/controllers"
	"github.com/pholguinc/api-go-matrices/internal/database"
	"github.com/pholguinc/api-go-matrices/internal/middlewares"
	"github.com/pholguinc/api-go-matrices/internal/repositories"
	"github.com/pholguinc/api-go-matrices/internal/routes"
	"github.com/pholguinc/api-go-matrices/internal/services"
)

// @title Matrix Factorization API
// @version 1.0
// @description API para realizar factorización QR de matrices rectangulares.
// @host localhost:3001
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Escribe 'Bearer ' seguido de tu token JWT.
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to Database
	database.ConnectDB()

	app := fiber.New()

	// CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// Middlewares
	app.Use(middlewares.Logger)

	// Environments
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// Dependency Injection - Repositories
	userRepo := repositories.NewUserRepository(database.DB)
	matrixRepo := repositories.NewMatrixRepository(database.DB)

	// Dependency Injection - Services
	authService := services.NewAuthService(userRepo)
	matrixService := services.NewMatrixService(matrixRepo)

	// Dependency Injection - Controllers
	authController := controllers.NewAuthController(authService)
	matrixController := controllers.NewMatrixController(matrixService)

	// Routes
	routes.SetupMatrixRoutes(app, matrixController)
	routes.SetupAuthRoutes(app, authController)

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

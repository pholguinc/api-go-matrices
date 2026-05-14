package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
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

var fiberLambda *fiberadapter.FiberLambdaV3

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
	// 1. Load environment variables (solo local)
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		_ = godotenv.Load()
	}

	// 2. Connect to Database
	database.ConnectDB()

	// 3. Setup App
	app := fiber.New()

	// Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	app.Use(middlewares.Logger)

	// Dependency Injection
	userRepo := repositories.NewUserRepository(database.DB)
	matrixRepo := repositories.NewMatrixRepository(database.DB)
	authService := services.NewAuthService(userRepo)
	matrixService := services.NewMatrixService(matrixRepo)
	authController := controllers.NewAuthController(authService)
	matrixController := controllers.NewMatrixController(matrixService)

	// Routes
	routes.SetupMatrixRoutes(app, matrixController)
	routes.SetupAuthRoutes(app, authController)

	// 4. Execution Mode
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		// Lambda Mode
		fiberLambda = fiberadapter.NewV3(app)
		lambda.Start(Handler)
	} else {
		// Local Mode
		port := os.Getenv("PORT")
		if port == "" {
			port = "3001"
		}
		log.Printf("Server starting locally on port %s", port)
		log.Fatal(app.Listen(":" + port))
	}
}

// Handler for AWS Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

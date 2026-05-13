package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Cargar puerto desde variable de entorno o usar 3001 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("API is running on port " + port)
	})

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

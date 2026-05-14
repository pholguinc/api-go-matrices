package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func Logger(c fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	stop := time.Now()

	log.Printf("[%s] %s %s - %v", c.Method(), c.Path(), stop.Sub(start), err)
	return err
}

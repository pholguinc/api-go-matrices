package database

import (
	"fmt"
	"log"
	"os"

	"github.com/pholguinc/api-go-matrices/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
			log.Fatal("Failed to connect to database: ", err)
		}
		return
	}

	log.Println("Connected to database successfully")

	DB.AutoMigrate(&models.User{}, &models.MatrixRecord{})
}

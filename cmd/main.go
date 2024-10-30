package main

import (
	"log"
	"player-wallet-api/config"
	"player-wallet-api/internal/delivery/http/router"
	"player-wallet-api/pkg/database"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize database
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Redis
	redisClient, err := database.NewRedisClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Echo
	e := echo.New()

	// Set up Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Setup routes
	router.SetupRoutes(e, db, redisClient)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

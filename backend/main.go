package main

import (
	"flag"
	"log"
	"os"
	"web3-portfolio-dashboard/backend/internal/api"
	"web3-portfolio-dashboard/backend/internal/config"
	"web3-portfolio-dashboard/backend/internal/database"
	"web3-portfolio-dashboard/backend/internal/services"

	"github.com/sirupsen/logrus"
)

func main() {
	// Parse command line flags
	testDB := flag.Bool("test-db", false, "Test database connection and exit")
	migrate := flag.Bool("migrate", false, "Run database migrations and exit")
	flag.Parse()

	// Load config from env
	cfg := config.Load()

	// Set up logger
	logger := logrus.New()
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// Connect to database
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test database connection
	if *testDB {
		log.Println("✅ Database connection test successful!")
		os.Exit(0)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Exit if only running migrations
	if *migrate {
		log.Println("✅ Database migrations completed successfully!")
		os.Exit(0)
	}

	// Create indexes for better performance
	if err := database.CreateIndexes(db); err != nil {
		log.Printf("Warning: Failed to create indexes: %v", err)
	}

	// Initialize services
	web3Service := services.NewWeb3Service(cfg)
	portfolioService := services.NewPortfolioService(db, web3Service)
	authService := services.NewAuthService(db, cfg.JWTSecret)
	alertService := services.NewAlertService(db)

	// Create and start the server
	server := api.NewServer(cfg, logger, db, portfolioService, authService, alertService, web3Service)
	if err := server.Start(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

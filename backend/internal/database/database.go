package database

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"web3-portfolio-dashboard/backend/internal/models"
)

// New creates a new database connection
func New(databaseURL string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Check if it's a SQLite connection (file path) or PostgreSQL URL
	if strings.HasPrefix(databaseURL, "sqlite://") || strings.HasSuffix(databaseURL, ".db") {
		// SQLite connection
		dbPath := strings.TrimPrefix(databaseURL, "sqlite://")
		if dbPath == "" {
			dbPath = "web3_portfolio.db"
		}

		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
		}
		log.Println("Connected to SQLite database:", dbPath)
	} else {
		// PostgreSQL connection
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
		}
		log.Println("Connected to PostgreSQL database")
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully")
	return db, nil
}

// Migrate runs database migrations
func Migrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Migrate all models at once
	if err := db.AutoMigrate(
		&models.User{},
		&models.Portfolio{},
		&models.Address{},
		&models.Transaction{},
		&models.Alert{},
		&models.Balance{},
		// Forum models
		&models.Question{},
		&models.Answer{},
		&models.Comment{},
		&models.Tag{},
		&models.Vote{},
		&models.Reputation{},
		&models.Role{},
	); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("✅ All tables migrated successfully")
	log.Println("Database migrations completed successfully")
	return nil
}

// CreateIndexes creates additional database indexes
func CreateIndexes(db *gorm.DB) error {
	log.Println("Creating database indexes...")

	// Create indexes for better performance
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);",
		"CREATE INDEX IF NOT EXISTS idx_portfolios_user_id ON portfolios(user_id);",
		"CREATE INDEX IF NOT EXISTS idx_addresses_portfolio_id ON addresses(portfolio_id);",
		"CREATE INDEX IF NOT EXISTS idx_addresses_network ON addresses(network);",
		"CREATE INDEX IF NOT EXISTS idx_transactions_portfolio_id ON transactions(portfolio_id);",
		"CREATE INDEX IF NOT EXISTS idx_transactions_tx_hash ON transactions(tx_hash);",
		"CREATE INDEX IF NOT EXISTS idx_transactions_network ON transactions(network);",
		"CREATE INDEX IF NOT EXISTS idx_transactions_timestamp ON transactions(timestamp);",
		"CREATE INDEX IF NOT EXISTS idx_alerts_user_id ON alerts(user_id);",
		"CREATE INDEX IF NOT EXISTS idx_alerts_type ON alerts(type);",
		"CREATE INDEX IF NOT EXISTS idx_balances_address_id ON balances(address_id);",
		"CREATE INDEX IF NOT EXISTS idx_balances_token_address ON balances(token_address);",
	}

	for _, index := range indexes {
		if err := db.Exec(index).Error; err != nil {
			return fmt.Errorf("failed to create index: %w", err)
		}
	}

	log.Println("Database indexes created successfully")
	return nil
}

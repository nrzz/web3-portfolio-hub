package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	RedisURL    string
	JWTSecret   string
	LogLevel    string
	Port        string

	// Web3 Configuration
	EthereumRPCURL string
	PolygonRPCURL  string
	BSCRPCURL      string
	ArbitrumRPCURL string

	// API Keys
	EtherscanAPIKey string
	CoinGeckoAPIKey string

	// Environment
	Environment string
}

func Load() *Config {
	// Load environment files in order of priority
	godotenv.Load(".env")
	godotenv.Load("env.dev")
	godotenv.Load("env.production")

	return &Config{
		DatabaseURL:     getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/web3_portfolio"),
		RedisURL:        getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		Port:            getEnv("PORT", "8080"),
		EthereumRPCURL:  getEnv("ETHEREUM_RPC_URL", "https://mainnet.infura.io/v3/your-project-id"),
		PolygonRPCURL:   getEnv("POLYGON_RPC_URL", ""),
		BSCRPCURL:       getEnv("BSC_RPC_URL", ""),
		ArbitrumRPCURL:  getEnv("ARBITRUM_RPC_URL", ""),
		EtherscanAPIKey: getEnv("ETHERSCAN_API_KEY", ""),
		CoinGeckoAPIKey: getEnv("COINGECKO_API_KEY", ""),
		Environment:     getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

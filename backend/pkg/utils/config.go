package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds application configuration
type Config struct {
	// Server configuration
	ServerPort string
	ServerHost string

	// Database configuration
	DatabaseURL string

	// Redis configuration
	RedisURL string

	// NATS configuration
	NATSURL     string
	NATSSubject string

	// API configuration
	CoinGeckoURL  string
	FetchInterval time.Duration
	FetchTimeout  time.Duration

	// Price filtering
	PriceChangeThreshold float64

	// Cache configuration
	CacheExpiration time.Duration

	// Blockchain configuration
	BlockchainRPCURL     string
	OracleContractAddr   string
	BlockchainPrivateKey string

	// Logging
	LogLevel string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		ServerHost:           getEnv("SERVER_HOST", "0.0.0.0"),
		DatabaseURL:          getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/oracle_db?sslmode=disable"),
		RedisURL:             getEnv("REDIS_URL", "redis://localhost:6379"),
		NATSURL:              getEnv("NATS_URL", "nats://localhost:4222"),
		NATSSubject:          getEnv("NATS_SUBJECT", "prices.ethusd"),
		CoinGeckoURL:         getEnv("COINGECKO_URL", "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"),
		FetchInterval:        getDurationEnv("FETCH_INTERVAL", "30s"),
		FetchTimeout:         getDurationEnv("FETCH_TIMEOUT", "10s"),
		PriceChangeThreshold: getFloatEnv("PRICE_CHANGE_THRESHOLD", 0.005), // 0.5%
		CacheExpiration:      getDurationEnv("CACHE_EXPIRATION", "1h"),
		BlockchainRPCURL:     getEnv("BLOCKCHAIN_RPC_URL", "http://localhost:8545"),
		OracleContractAddr:   getEnv("ORACLE_CONTRACT_ADDR", "0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		BlockchainPrivateKey: getEnv("BLOCKCHAIN_PRIVATE_KEY", ""),
		LogLevel:             getEnv("LOG_LEVEL", "info"),
	}

	// Validate required configurations
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.ServerPort == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}
	if c.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	if c.RedisURL == "" {
		return fmt.Errorf("REDIS_URL is required")
	}
	if c.NATSURL == "" {
		return fmt.Errorf("NATS_URL is required")
	}
	if c.CoinGeckoURL == "" {
		return fmt.Errorf("COINGECKO_URL is required")
	}
	if c.PriceChangeThreshold < 0 || c.PriceChangeThreshold > 1 {
		return fmt.Errorf("PRICE_CHANGE_THRESHOLD must be between 0 and 1")
	}
	return nil
}

// GetServerAddr returns the server address
func (c *Config) GetServerAddr() string {
	return c.ServerHost + ":" + c.ServerPort
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getDurationEnv gets a duration environment variable with a default value
func getDurationEnv(key, defaultValue string) time.Duration {
	value := getEnv(key, defaultValue)
	duration, err := time.ParseDuration(value)
	if err != nil {
		// Return default if parsing fails
		duration, _ = time.ParseDuration(defaultValue)
	}
	return duration
}

// getFloatEnv gets a float environment variable with a default value
func getFloatEnv(key string, defaultValue float64) float64 {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}

	return floatValue
}

// getIntEnv gets an integer environment variable with a default value
func getIntEnv(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}

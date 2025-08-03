package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port     string
	LogLevel string

	// NewRelic
	NewRelicLicenseKey string
	NewRelicAppName    string

	// Trading
	DryRun           bool
	BinanceAPIKey    string
	BinanceSecretKey string

	// Arbitrage
	MinProfitThreshold float64
	MaxTradeAmount     float64
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or could not be loaded: %v", err)
	}

	dryRun, err := strconv.ParseBool(getEnv("DRY_RUN", "true"))
	if err != nil {
		log.Fatalf("Failed to parse DRY_RUN: %v", err)
	}

	minProfitThreshold, err := strconv.ParseFloat(getEnv("MIN_PROFIT_THRESHOLD", "0.5"), 64)
	if err != nil {
		log.Fatalf("Failed to parse MIN_PROFIT_THRESHOLD: %v", err)
	}
	maxTradeAmount, err := strconv.ParseFloat(getEnv("MAX_TRADE_AMOUNT", "1000"), 64)
	if err != nil {
		log.Fatalf("Failed to parse MAX_TRADE_AMOUNT: %v", err)
	}

	return &Config{
		Port:     getEnv("PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "INFO"),

		NewRelicLicenseKey: getEnv("NEW_RELIC_LICENSE_KEY", ""),
		NewRelicAppName:    getEnv("NEW_RELIC_APP_NAME", "delta-bot"),

		DryRun:           dryRun,
		BinanceAPIKey:    getEnv("BINANCE_API_KEY", ""),
		BinanceSecretKey: getEnv("BINANCE_SECRET_KEY", ""),

		MinProfitThreshold: minProfitThreshold,
		MaxTradeAmount:     maxTradeAmount,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

package exchange

import (
	"context"

	"github.com/haydenfowler/delta-bot/internal/logger"
)

// Exchange defines a minimal interface for any exchange
type Exchange interface {
	// Name returns the exchange name
	Name() string

	// Start begins the exchange operations (connect, monitor, trade)
	Start(ctx context.Context) error

	// Stop gracefully shuts down the exchange
	Stop(ctx context.Context) error

	// IsRunning returns whether the exchange is currently active
	IsRunning() bool
}

// ExchangeType represents the type of exchange
type ExchangeType string

const (
	ExchangeTypeBinance  ExchangeType = "binance"
	ExchangeTypeKucoin   ExchangeType = "kucoin"
	ExchangeTypeCoinbase ExchangeType = "coinbase"
	ExchangeTypeKraken   ExchangeType = "kraken"
)

// Config holds exchange configuration
type Config struct {
	Type    ExchangeType `json:"type"`
	Name    string       `json:"name"`
	APIKey  string       `json:"api_key"`
	Secret  string       `json:"secret"`
	Sandbox bool         `json:"sandbox"`
	Enabled bool         `json:"enabled"`
}

// Factory creates exchange instances
type Factory struct {
	logger *logger.Logger
}

// NewFactory creates a new exchange factory
func NewFactory(logger *logger.Logger) *Factory {
	return &Factory{logger: logger}
}

// Create creates an exchange instance
func (f *Factory) Create(config Config) (Exchange, error) {
	switch config.Type {
	case ExchangeTypeBinance:
		return NewBinance(config, f.logger), nil
	case ExchangeTypeKucoin:
		return NewKucoin(config, f.logger), nil
	case ExchangeTypeCoinbase:
		return NewCoinbase(config, f.logger), nil
	case ExchangeTypeKraken:
		return NewKraken(config, f.logger), nil
	default:
		return nil, NewUnsupportedExchangeError(config.Type)
	}
}

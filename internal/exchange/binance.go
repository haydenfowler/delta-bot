package exchange

import (
	"context"
	"fmt"

	"github.com/haydenfowler/delta-bot/internal/logger"
)

// Binance implements the Exchange interface for Binance
type Binance struct {
	config  Config
	logger  *logger.Logger
	running bool
}

// NewBinance creates a new Binance exchange instance
func NewBinance(config Config, logger *logger.Logger) *Binance {
	return &Binance{
		config: config,
		logger: logger,
	}
}

// Name returns the exchange name
func (b *Binance) Name() string {
	return b.config.Name
}

// Start begins Binance operations
func (b *Binance) Start(ctx context.Context) error {
	if b.running {
		return fmt.Errorf("binance exchange already running")
	}

	b.logger.Info(fmt.Sprintf("Starting %s exchange", b.Name()))

	// Here Binance would:
	// 1. Connect to Binance API
	// 2. Start monitoring price feeds
	// 3. Detect arbitrage opportunities
	// 4. Execute trades when profitable

	b.running = true
	b.logger.Info(fmt.Sprintf("%s exchange started successfully", b.Name()))

	return nil
}

// Stop gracefully shuts down Binance operations
func (b *Binance) Stop(ctx context.Context) error {
	if !b.running {
		return nil
	}

	b.logger.Info(fmt.Sprintf("Stopping %s exchange", b.Name()))

	// Here Binance would:
	// 1. Cancel any pending orders
	// 2. Close WebSocket connections
	// 3. Clean up resources

	b.running = false
	b.logger.Info(fmt.Sprintf("%s exchange stopped", b.Name()))

	return nil
}

// IsRunning returns whether the exchange is currently active
func (b *Binance) IsRunning() bool {
	return b.running
}

package exchange

import (
	"context"
	"fmt"

	"github.com/haydenfowler/delta-bot/internal/logger"
)

// Coinbase implements the Exchange interface for Coinbase
type Coinbase struct {
	config  Config
	logger  *logger.Logger
	running bool
}

// NewCoinbase creates a new Coinbase exchange instance
func NewCoinbase(config Config, logger *logger.Logger) *Coinbase {
	return &Coinbase{
		config: config,
		logger: logger,
	}
}

// Name returns the exchange name
func (c *Coinbase) Name() string {
	return c.config.Name
}

// Start begins Coinbase operations
func (c *Coinbase) Start(ctx context.Context) error {
	if c.running {
		return fmt.Errorf("coinbase exchange already running")
	}

	c.logger.Info(fmt.Sprintf("Starting %s exchange", c.Name()))
	c.running = true
	c.logger.Info(fmt.Sprintf("%s exchange started successfully", c.Name()))

	return nil
}

// Stop gracefully shuts down Coinbase operations
func (c *Coinbase) Stop(ctx context.Context) error {
	if !c.running {
		return nil
	}

	c.logger.Info(fmt.Sprintf("Stopping %s exchange", c.Name()))
	c.running = false
	c.logger.Info(fmt.Sprintf("%s exchange stopped", c.Name()))

	return nil
}

// IsRunning returns whether the exchange is currently active
func (c *Coinbase) IsRunning() bool {
	return c.running
}

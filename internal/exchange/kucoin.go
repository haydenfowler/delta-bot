package exchange

import (
	"context"
	"fmt"

	"github.com/haydenfowler/delta-bot/internal/logger"
)

// Kucoin implements the Exchange interface for KuCoin
type Kucoin struct {
	config  Config
	logger  *logger.Logger
	running bool
}

// NewKucoin creates a new KuCoin exchange instance
func NewKucoin(config Config, logger *logger.Logger) *Kucoin {
	return &Kucoin{
		config: config,
		logger: logger,
	}
}

// Name returns the exchange name
func (k *Kucoin) Name() string {
	return k.config.Name
}

// Start begins KuCoin operations
func (k *Kucoin) Start(ctx context.Context) error {
	if k.running {
		return fmt.Errorf("kucoin exchange already running")
	}

	k.logger.Info(fmt.Sprintf("Starting %s exchange", k.Name()))
	k.running = true
	k.logger.Info(fmt.Sprintf("%s exchange started successfully", k.Name()))

	return nil
}

// Stop gracefully shuts down KuCoin operations
func (k *Kucoin) Stop(ctx context.Context) error {
	if !k.running {
		return nil
	}

	k.logger.Info(fmt.Sprintf("Stopping %s exchange", k.Name()))
	k.running = false
	k.logger.Info(fmt.Sprintf("%s exchange stopped", k.Name()))

	return nil
}

// IsRunning returns whether the exchange is currently active
func (k *Kucoin) IsRunning() bool {
	return k.running
}

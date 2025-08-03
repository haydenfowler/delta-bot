package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/haydenfowler/delta-bot/internal/config"
	"github.com/haydenfowler/delta-bot/internal/logger"
	"github.com/haydenfowler/delta-bot/internal/server"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger with NewRelic
	appLogger := logger.New(cfg.NewRelicLicenseKey, cfg.NewRelicAppName)
	defer appLogger.Shutdown()

	appLogger.Info("Delta Bot starting up")

	// Create and start server
	srv := server.New(cfg.Port, appLogger)

	// Channel to listen for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := srv.Start(cfg.Port); err != nil {
			appLogger.Error("Server failed to start", err)
			os.Exit(1)
		}
	}()

	appLogger.Info("Server started successfully. Press Ctrl+C to shutdown.")

	// Wait for interrupt signal
	<-quit
	appLogger.Info("Shutdown signal received")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		appLogger.Error("Server forced to shutdown", err)
		os.Exit(1)
	}

	appLogger.Info("Server exited")
}

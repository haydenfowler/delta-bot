package routes

import (
	"github.com/haydenfowler/delta-bot/internal/handlers"
	"github.com/haydenfowler/delta-bot/internal/logger"
	"github.com/haydenfowler/delta-bot/internal/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SetupRoutes configures all application routes and middleware
func SetupRoutes(e *echo.Echo, logger *logger.Logger, nrApp *newrelic.Application) {
	// Basic middleware
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())

	// Custom middleware
	e.Use(middleware.NewRelicMiddleware(nrApp))
	e.Use(middleware.RequestLogger(logger))

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(logger)

	// Health routes
	e.GET("/health", healthHandler.Health)
}

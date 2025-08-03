package server

import (
	"context"
	"fmt"
	"time"

	"github.com/haydenfowler/delta-bot/internal/logger"
	"github.com/haydenfowler/delta-bot/internal/routes"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo   *echo.Echo
	logger *logger.Logger
}

func New(port string, logger *logger.Logger) *Server {
	e := echo.New()

	// Hide Echo banner
	e.HideBanner = true

	// Configure timeouts
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 30 * time.Second

	s := &Server{
		echo:   e,
		logger: logger,
	}

	// Setup all routes and middleware
	routes.SetupRoutes(e, logger, logger.GetNewRelicApp())

	return s
}

func (s *Server) Start(port string) error {
	s.logger.Info(fmt.Sprintf("Starting Echo server on port %s", port))
	return s.echo.Start(":" + port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down Echo server")
	return s.echo.Shutdown(ctx)
}

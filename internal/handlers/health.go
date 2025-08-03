package handlers

import (
	"net/http"
	"time"

	"github.com/haydenfowler/delta-bot/internal/logger"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	logger *logger.Logger
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHealthHandler(logger *logger.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger,
	}
}

// Health handles GET /health requests
func (h *HealthHandler) Health(c echo.Context) error {
	h.logger.Info("Health check requested")
	h.logger.SendHeartbeat()

	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
	}

	return c.JSON(http.StatusOK, response)
}

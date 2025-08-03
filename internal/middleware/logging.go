package middleware

import (
	"fmt"

	"github.com/haydenfowler/delta-bot/internal/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RequestLogger returns Echo middleware for request logging
func RequestLogger(logger *logger.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		LogMethod:   true,
		LogLatency:  true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.Info(formatRequestLog(v))
			} else {
				logger.Error(formatRequestLog(v), v.Error)
			}
			return nil
		},
	})
}

func formatRequestLog(v middleware.RequestLoggerValues) string {
	return fmt.Sprintf("HTTP %d %s %s - %v",
		v.Status, v.Method, v.URI, v.Latency)
}

package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// NewRelicMiddleware returns Echo middleware for NewRelic integration
func NewRelicMiddleware(app *newrelic.Application) echo.MiddlewareFunc {
	if app == nil {
		// Return a no-op middleware if NewRelic is not configured
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return next
		}
	}
	return nrecho.Middleware(app)
}

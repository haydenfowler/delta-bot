package logger

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type Logger struct {
	nrApp *newrelic.Application
}

func New(licenseKey, appName string) *Logger {
	var nrApp *newrelic.Application
	var err error

	if licenseKey != "" {
		nrApp, err = newrelic.NewApplication(
			newrelic.ConfigAppName(appName),
			newrelic.ConfigLicense(licenseKey),
			newrelic.ConfigInfoLogger(os.Stdout),
		)
		if err != nil {
			log.Printf("Failed to initialize NewRelic: %v", err)
			nrApp = nil
		} else {
			log.Println("NewRelic initialized successfully")
		}
	}

	return &Logger{
		nrApp: nrApp,
	}
}

func (l *Logger) Info(message string) {
	log.Printf("INFO: %s", message)
}

func (l *Logger) Error(message string, err error) {
	log.Printf("ERROR: %s: %v", message, err)
	if l.nrApp != nil {
		l.nrApp.RecordCustomEvent("ErrorEvent", map[string]interface{}{
			"message": message,
			"error":   err.Error(),
		})
	}
}

func (l *Logger) RecordEvent(eventType string, attributes map[string]interface{}) {
	if l.nrApp != nil {
		l.nrApp.RecordCustomEvent(eventType, attributes)
	}
}

func (l *Logger) SendHeartbeat() {
	l.Info("Sending heartbeat")
	if l.nrApp != nil {
		l.RecordEvent("HeartbeatEvent", map[string]interface{}{
			"type":   "heartbeat",
			"status": "healthy",
		})
	}
}

func (l *Logger) Shutdown() {
	if l.nrApp != nil {
		l.nrApp.Shutdown(0)
	}
}

// GetNewRelicApp returns the NewRelic application instance for middleware use
func (l *Logger) GetNewRelicApp() *newrelic.Application {
	return l.nrApp
}

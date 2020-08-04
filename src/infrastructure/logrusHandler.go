package infrastructure

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger defines the logger
type Logger struct {
	logger *logrus.Logger
}

// NewLogger creates a new logger
func NewLogger() Logger {
	var logger = Logger{
		logrus.New(),
	}

	logger.logger.Out = os.Stdout
	logger.logger.SetReportCaller(true)
	logger.logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}

// Info logs an info event
func (logger Logger) Info(message string) {
	logger.logger.Info(message)
}

// Debug logs a debug event
func (logger Logger) Debug(message string) {
	logger.logger.Debug(message)

}

// Error logs an error event
func (logger Logger) Error(message string) {
	logger.logger.Error(message)
}

// Fatal logs a fatal event
func (logger Logger) Fatal(message string) {
	logger.logger.Fatal(message)
}

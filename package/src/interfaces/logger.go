package interfaces

// LoggerInterface defines the interface for logger
type LoggerInterface interface {
	Fatal(message string)
	Error(message string)
	Debug(message string)
	Info(message string)
}

// CustomLogger defines my logger
type CustomLogger struct {
	logger LoggerInterface
}

// Logger is the type of my custom logger
type Logger CustomLogger

// Info logs an "info" event
func (customLogger *CustomLogger) Info(message string) {
	customLogger.logger.Info(message)
}

// Debug logs a "debug" event
func (customLogger *CustomLogger) Debug(message string) {
	customLogger.logger.Debug(message)
}

// Error logs an "error" event
func (customLogger *CustomLogger) Error(message string) {
	customLogger.logger.Error(message)
}

// Fatal logs a "fatal" event
func (customLogger *CustomLogger) Fatal(message string) {
	customLogger.logger.Fatal(message)
}

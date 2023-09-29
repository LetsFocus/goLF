package logger

import (
	"time"

	"github.com/fatih/color"
)

// CustomLogger represents our custom logger.
type CustomLogger struct {
	infoColor    *color.Color
	warningColor *color.Color
	errorColor   *color.Color
}

// NewCustomLogger creates a new instance of CustomLogger.
func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		infoColor:    color.New(color.FgCyan),
		warningColor: color.New(color.FgYellow),
		errorColor:   color.New(color.FgRed),
	}
}

// Info logs an info message in cyan color.
func (cl *CustomLogger) Info(message string) {
	cl.infoColor.Printf("[INFO] %v  %s\n", time.Now().Format("2006/01/02 - 15:04:05"), message)
}

// Warning logs a warning message in yellow color.
func (cl *CustomLogger) Warning(message string) {
	cl.warningColor.Printf("[WARN] %v %s\n", time.Now().Format("2006/01/02 - 15:04:05"), message)
}

// Error logs an error message in red color.
func (cl *CustomLogger) Error(message string) {
	cl.errorColor.Printf("[ERROR] %v %s\n", time.Now().Format("2006/01/02 - 15:04:05"), message)
}

// Infof logs an info message with formatted arguments.
func (cl *CustomLogger) Infof(format string, args ...interface{}) {
	timeAppendedArgs := append([]interface{}{time.Now().Format("2006/01/02 - 15:04:05")}, args...)
	cl.infoColor.Printf("[INFO] %v "+format+"\n", timeAppendedArgs...)
}

// Errorf logs an error message with formatted arguments.
func (cl *CustomLogger) Errorf(format string, args ...interface{}) {
	timeAppendedArgs := append([]interface{}{time.Now().Format("2006/01/02 - 15:04:05")}, args...)
	cl.errorColor.Printf("[ERROR] %v "+format+"\n", timeAppendedArgs...)
}

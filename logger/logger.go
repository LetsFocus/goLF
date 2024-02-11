package logger

import (
	"bytes"
	"fmt"
	"time"

	"github.com/fatih/color"
)

// CustomLogger represents our custom logger.
type CustomLogger struct {
	infoColor    *color.Color
	warningColor *color.Color
	errorColor   *color.Color
	buffer       *bytes.Buffer
}

// NewCustomLogger creates a new instance of CustomLogger.
func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		infoColor:    color.New(color.FgCyan),
		warningColor: color.New(color.FgYellow),
		errorColor:   color.New(color.FgRed),
		buffer:       bytes.NewBuffer(nil),
	}
}

// Info logs an info message in cyan color.
func (cl *CustomLogger) Info(message string) {
	logEntry := fmt.Sprintf("[INFO] %v  %s\n", time.Now().Format("2006/01/02 - 15:04:05"), message)
	cl.infoColor.Print(logEntry)
	cl.buffer.WriteString(logEntry)
}

// Warning logs a warning message in yellow color.
func (cl *CustomLogger) Warning(message string) {
	logEntry := fmt.Sprintf("[WARN] %v  %s\n", time.Now().Format("2006/01/02 - 15:04:05"), message)
	cl.warningColor.Print(logEntry)
	cl.buffer.WriteString(logEntry)
}

// Error logs an error message in red color.
func (cl *CustomLogger) Error(message string) {
	logEntry := fmt.Sprintf("[ERROR] %v  %s\n", time.Now().Format("2006/01/02 - 15:04:05"), message)
	cl.errorColor.Print(logEntry)
	cl.buffer.WriteString(logEntry)
}

// Infof logs an info message with formatted arguments.
func (cl *CustomLogger) Infof(format string, args ...interface{}) {
	timeAppendedArgs := append([]interface{}{time.Now().Format("2006/01/02 - 15:04:05")}, args...)
	logEntry := fmt.Sprintf("[INFO] %v "+format+"\n", timeAppendedArgs...)
	cl.infoColor.Printf(logEntry)
	cl.buffer.WriteString(logEntry)
}

// Errorf logs an error message with formatted arguments.
func (cl *CustomLogger) Errorf(format string, args ...interface{}) {
	timeAppendedArgs := append([]interface{}{time.Now().Format("2006/01/02 - 15:04:05")}, args...)
	logEntry := fmt.Sprintf("[ERROR] %v "+format+"\n", timeAppendedArgs...)
	cl.errorColor.Printf(logEntry)
	cl.buffer.WriteString(logEntry)
}

func (cl *CustomLogger) GetLog() string {
	return cl.buffer.String()
}

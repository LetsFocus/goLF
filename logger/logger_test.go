package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomLogger_LogMethods(t *testing.T) {
	tests := []struct {
		name           string
		logMethod      func(logger *CustomLogger, message string)
		expectedOutput string
	}{
		{
			name:           "Info",
			logMethod:      (*CustomLogger).Info,
			expectedOutput: "[INFO]",
		},
		{
			name:           "Warning",
			logMethod:      (*CustomLogger).Warning,
			expectedOutput: "[WARN]",
		},
		{
			name:           "Error",
			logMethod:      (*CustomLogger).Error,
			expectedOutput: "[ERROR]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewCustomLogger()

			// Call the log method with a message
			message := "Test message"
			tt.logMethod(logger, message)

			// Check if the log contains the expected output
			expectedLogEntry := fmt.Sprintf("%s %v  %s\n", tt.expectedOutput, time.Now().Format("2006/01/02 - 15:04:05"), message)
			assert.Contains(t, logger.GetLog(), expectedLogEntry)
		})
	}
}

func TestCustomLogger_FormattedLogMethods(t *testing.T) {
	tests := []struct {
		name           string
		logMethod      func(logger *CustomLogger, format string, args ...interface{})
		expectedOutput string
	}{
		{
			name:           "Infof",
			logMethod:      (*CustomLogger).Infof,
			expectedOutput: "[INFO]",
		},
		{
			name:           "Errorf",
			logMethod:      (*CustomLogger).Errorf,
			expectedOutput: "[ERROR]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewCustomLogger()

			// Call the formatted log method with a format and arguments
			format := "Formatted message with %s."
			arg := "args"
			tt.logMethod(logger, format, arg)

			// Check if the log contains the expected output
			expectedLogEntry := fmt.Sprintf("%s %v "+format+"\n", tt.expectedOutput, time.Now().Format("2006/01/02 - 15:04:05"), arg)
			assert.Contains(t, logger.GetLog(), expectedLogEntry)
		})
	}
}

func TestCustomLogger_GetLog(t *testing.T) {
	tests := []struct {
		name           string
		logMethods     func(logger *CustomLogger)
		expectedOutput string
	}{
		{
			name: "Info and Warning",
			logMethods: func(logger *CustomLogger) {
				logger.Info("Info message")
				logger.Warning("Warning message")
			},
			expectedOutput: "[INFO] " + time.Now().Format("2006/01/02 - 15:04:05") + "  Info message\n" +
				"[WARN] " + time.Now().Format("2006/01/02 - 15:04:05") + "  Warning message\n",
		},
		{
			name: "Errorf",
			logMethods: func(logger *CustomLogger) {
				format := "Formatted error message with %s."
				arg := "args"
				logger.Errorf(format, arg)
			},
			expectedOutput: "[ERROR] " + time.Now().Format("2006/01/02 - 15:04:05") + " Formatted error message with args.\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewCustomLogger()

			// Call the log methods
			tt.logMethods(logger)

			// Check if the log contains the expected output
			assert.Equal(t, tt.expectedOutput, logger.GetLog())
		})
	}
}

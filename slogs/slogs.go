package slogs

import (
	"log/slog"
	"os"
	"time"
)

const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
	WARN  = "WARN"
	ERROR = "ERROR"
)

type Log struct {
	Logger *slog.Logger
	level  *slog.LevelVar
}

func NewLogger() Log {
	logLevel := &slog.LevelVar{}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)

	log := Log{Logger: logger, level: logLevel}

	go watchLogLevelChanges(log)

	return log
}

func watchLogLevelChanges(l Log) {
	for {
		currentLogLevel := os.Getenv("LOG_LEVEL")
		l.Logger.Info(currentLogLevel)
		if currentLogLevel != "" && currentLogLevel != l.level.Level().String() {
			l.level.Set(getSlogLevel(currentLogLevel))
			l.Logger.Info("Log level updated", "level", currentLogLevel)
		}

		time.Sleep(time.Second * 1)
	}
}

func getSlogLevel(level string) slog.Level {
	switch level {
	case DEBUG:
		return slog.LevelDebug
	case ERROR:
		return slog.LevelError
	case WARN:
		return slog.LevelWarn
	default:
		return slog.LevelInfo
	}
}

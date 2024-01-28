package configs

import (
	"github.com/LetsFocus/goLF/logger"
	"os"

	"github.com/joho/godotenv"
)

func NewConfig(log *logger.CustomLogger) Config {
	env := os.Getenv("APP_ENV")
	envPath := ""
	if env != "" {
		envPath = "./configs/." + env + ".env"
	} else {
		envPath = "./configs/.env"
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Error("No .env file found")
	}

	log.Info("Logs are initialized")
	return Config{Log: log}
}

type Config struct {
	Log *logger.CustomLogger
}

type Configs interface {
	Get(key string) string
}

func (c Config) Get(key string) string {
	return os.Getenv(key)
}

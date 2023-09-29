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
		envPath = "./config/." + env + ".env"
	} else {
		envPath = "./config/.env"
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
	GetorDefault(key, defaultVal string) string
}

func (c Config) Get(key string) string {
	return os.Getenv(key)
}

func (c Config) GetorDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}

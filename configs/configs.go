package configs

import (
	"github.com/LetsFocus/goLF/slogs"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func NewConfig(log slogs.Log) Config {
	loadConfigs(log)

	if configsRefresh := os.Getenv("CONFIG_REFRESH"); configsRefresh == "true" {
		go watchAndReloadConfigs(log)

	}

	log.Logger.Info("configs are loaded")

	return Config{Log: log}
}

type Config struct {
	Log slogs.Log
}

type Configs interface {
	Get(key string) string
}

func (c Config) Get(key string) string {
	return os.Getenv(key)
}

func watchAndReloadConfigs(log slogs.Log) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			loadConfigs(log)
			log.Logger.Info("counter")
		}
	}
}

func loadConfigs(log slogs.Log) {
	env := os.Getenv("APP_ENV")
	envPath := ""
	if env != "" {
		envPath = "./configs/." + env + ".env"
	} else {
		envPath = "./configs/.env"
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Logger.Error("No .env file found")
	}
}

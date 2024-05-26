package configs

import (
	"os"
	"path/filepath"

	"github.com/LetsFocus/goLF/logger"
	"github.com/joho/godotenv"
)

func findConfigsDir(dir string) (string, error) {
	depth := 0
	maxDepth := 3
	configsDir := filepath.Join(dir, "configs")
	if _, err := os.Stat(configsDir); err == nil {
		return configsDir, nil
	}
	if depth < maxDepth {
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", os.ErrNotExist
		}
		depth++
		findConfigsDir(parentDir)
	}
	return "", os.ErrNotExist
}

func NewConfig(log *logger.CustomLogger) Config {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Error("Unable to get current directory")
		return Config{Log: log}
	}
	path, err := findConfigsDir(currentDir)
	if err != nil {
		log.Infof("No configs directory found:%v", err)
	} else {

		envPath := filepath.Join(path, ".env")
		if err := godotenv.Load(envPath); err != nil {
			log.Error("No .env file found")
		}
		env := os.Getenv("APP_ENV")
		appEnvPath := ""
		if env != "" {
			appEnvPath = filepath.Join(path, env, ".env")
			if err := godotenv.Load(appEnvPath); err != nil {
				log.Error("No app_env file found")
			}
		}
		log.Infof("Logs are initialized path: %v", envPath)
	}
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

package configs

import (
	"github.com/LetsFocus/goLF/errors"
	"os"
	"path/filepath"

	"github.com/LetsFocus/goLF/logger"
	"github.com/joho/godotenv"
)

// NewConfig initializes and returns a new Config instance.
func NewConfig(log *logger.CustomLogger) *Config {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Error("Unable to get current directory")
		return &Config{}
	}
	path, err := findConfigsDir(currentDir)
	if err != nil {
		log.Infof("No configs directory found:%v", err)
		return &Config{}
	}
	envPath := filepath.Join(path, ".env")
	if !loadEnv(log, envPath) {
		return &Config{}
	}
	env := os.Getenv("APP_ENV")
	appEnvPath := ""
	if env != "" {
		appEnvPath = filepath.Join(path, env, ".env")
		if !loadEnv(log, appEnvPath) {
			return &Config{}
		}
	}
	log.Infof("Logs are initialized path: %v", envPath)
	return &Config{Log: log}
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

func (c Config) GetOrDefault(key, defaultValue string) string {
	if key == "" {
		return defaultValue
	}
	if os.Getenv(key) == "" {
		return defaultValue
	}
	return c.Get(key)
}

// findConfigsDir searches for a "configs" directory by traversing up from the specified directory.
func findConfigsDir(dir string) (string, error) {
	for i := 0; i < MaxParentSearchDepth; i++ {
		configsDir := filepath.Join(dir, "configs")
		if _, err := os.Stat(configsDir); err == nil {
			return configsDir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", &errors.MissingDir{Param: "Configs"}
		}
		dir = parentDir
	}
	return "", &errors.MissingDir{Param: "Configs"}
}

// loadEnv attempts to load environment variables from a specified .env file.
func loadEnv(log *logger.CustomLogger, envPath string) bool {
	if err := godotenv.Load(envPath); err != nil {
		log.Errorf("Unable to load .env file at path: %s", envPath)
		return false
	}
	return true
}

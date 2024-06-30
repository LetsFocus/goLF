package configs

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"github.com/LetsFocus/goLF/logger"

	"github.com/LetsFocus/goLF/errors"
)

// Config holds the configuration settings.
type Config struct {
	Log        *logger.CustomLogger
	currentDir string
	path       string
	envPath    string
}

// NewConfig initializes and returns a new Config instance.
func NewConfig(log *logger.CustomLogger) *Config {
	config := &Config{Log: log}

	if err := config.setCurrentDirectory(); err != nil {
		log.Error("Unable to get current directory")
		return config
	}

	if err := config.setConfigPath(); err != nil {
		log.Errorf("No configs directory found: %v", err)
		return config
	}

	if err := config.loadEnvironmentVariables(); err != nil {
		log.Errorf("Failed to load: %v", err)
		return config
	}

	log.Infof("Logs are initialized path: %v", config.envPath)
	return config
}

// setCurrentDirectory sets the current directory in the config.
func (c *Config) setCurrentDirectory() error {
	var err error
	c.currentDir, err = os.Getwd()
	return err

}

// setConfigPath sets the path to the configuration directory.
func (c *Config) setConfigPath() error {
	var err error
	c.path, err = findConfigsDir(c.currentDir)
	return err

}

// loadEnvironmentVariables loads the environment variables from .env files.
func (c *Config) loadEnvironmentVariables() error {
	if err := loadEnv(c.Log, filepath.Join(c.path, ".env")); err != nil {
		return err
	}

	env := os.Getenv("APP_ENV")
	if env != "" {
		if err := loadEnv(c.Log, filepath.Join(c.path, env, ".env")); err != nil {
			return err
		}
	}

	c.envPath = filepath.Join(c.path, ".env")
	return nil
}

// Get retrieves the value of an environment variable given its key.
func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

// GetOrDefault retrieves the value of an environment variable given its key, or returns a default value if the environment variable is not set or if the key is empty.
func (c *Config) GetOrDefault(key, defaultValue string) string {
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
			return "", errors.MissingDir{Param: "Configs"}
		}

		dir = parentDir
	}
	return "", errors.MissingDir{Param: "Configs"}
}

// loadEnv attempts to load environment variables from a specified .env file.
func loadEnv(log *logger.CustomLogger, envPath string) error {
	if err := godotenv.Load(envPath); err != nil {
		log.Errorf("Unable to load .env file at path: %s", envPath)
		return err
	}

	return nil
}

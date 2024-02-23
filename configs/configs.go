package configs

import (
	"os"

	"github.com/LetsFocus/goLF/slogs"
	"github.com/joho/godotenv"
)

func NewConfig(log slogs.Log, path string) Config {
	c := Config{log: log}
	c.LoadConfigs(log, path)

	log.Logger.Info("configs are loaded")

	return c
}

type Config struct {
	log  slogs.Log
	path []string
}

type Configs interface {
	Get(key string) string
	GetPath() string
}

func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

func (c *Config) GetPath() []string {
	return c.path
}

func (c *Config) LoadConfigs(log slogs.Log, path string) {
	env := os.Getenv("APP_ENV")

	envPath := make([]string, 0)

	location := getLocation()

	if path != "" {
		envPath = append(envPath, path+"/.env")
	}

	if env != "" {
		envPath = append(envPath, location+".env")
		envPath = append(envPath, location+".local.env")
		envPath = append(envPath, location+env+".env")
	} else {
		envPath = append(envPath, location+".env")
		envPath = append(envPath, location+".local.env")
	}

	if err := godotenv.Load(envPath...); err != nil {
		log.Logger.Error("No .env file found")
	}

	c.path = envPath
}

func getLocation() string {
	defaultLocations := make([]string, 0)
	defaultLocations = append(defaultLocations, "./configs/", "../configs/", "../../configs/")

	for _, path := range defaultLocations {
		if _, err := os.Stat(path); err != nil {
			continue
		}

		return path
	}

	return ""
}

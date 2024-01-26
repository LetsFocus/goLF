package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/LetsFocus/goLF/logger"
)

func TestNewConfig_LoadsConfigWithEnv(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	os.Setenv("TEST_KEY", "test_value")
	
	logger := logger.NewCustomLogger()
	config := NewConfig(logger)

	assert.NotNil(t, config.Log)
	assert.Equal(t, "test_value", config.Get("TEST_KEY"))

	os.Unsetenv("APP_ENV")
	os.Unsetenv("TEST_KEY")
}

func TestNewConfig_LoadsConfigWithoutEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	
	logger := logger.NewCustomLogger()
	config := NewConfig(logger)
	
	assert.NotNil(t, config.Log)
	assert.Equal(t, "test_value", config.Get("TEST_KEY"))

	os.Unsetenv("TEST_KEY")
}
package database

import (
	"errors"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_InitializeRedis(t *testing.T) {
	t.Setenv("REDIS_LOCALHOST", "127.0.0.1")
	t.Setenv("REDIS_PORT", "6379")
	testcases := []struct {
		input *model.GoLF
	}{
		{
			input: &model.GoLF{
				Config: configs.Config{Log: logger.NewCustomLogger()},
				Logger: logger.NewCustomLogger(),
			},
		},
	}

	for _, tc := range testcases {
		InitializeRedis(tc.input, "")
	}
}

func Test_createRedisConnectionFail(t *testing.T) {
	log := logger.NewCustomLogger()

	testcases := []struct {
		desc        string
		redisConfig redisConfig
		err         error
	}{

		{
			desc: "Failed to connect",
			redisConfig: redisConfig{
				addr:           "127.0.0.1:6399",
				password:       "",
				dB:             0,
				retries:        5,
				retryTime:      time.Duration(5),
				poolSize:       10,
				minIdleConns:   4,
				maxIdleConns:   8,
				connMaxLife:    time.Duration(10),
				conMaxIdleTime: time.Duration(30),
			},
			err: errors.New("dial tcp 127.0.0.1:6399: connectex: No connection could be made because the target machine actively refused it."),
		},
	}
	for i, tc := range testcases {
		_, err := createRedisConnection(&tc.redisConfig, log)
		assert.Equalf(t, tc.err.Error(), err.Error(), "Test[%d] FAILED, Could not connect to SQL, got error: %v\n", i, err)
	}
}
func Test_createRedisConnectionPass(t *testing.T) {
	log := logger.NewCustomLogger()

	testcases := []struct {
		desc        string
		redisConfig redisConfig
		err         error
	}{
		{
			desc: "Successfully connected",
			redisConfig: redisConfig{
				addr:           "127.0.0.1:6379",
				password:       "",
				dB:             0,
				retries:        5,
				retryTime:      time.Duration(5),
				poolSize:       10,
				minIdleConns:   4,
				maxIdleConns:   8,
				connMaxLife:    time.Duration(10),
				conMaxIdleTime: time.Duration(30),
			},
			err: nil,
		},
	}
	for i, tc := range testcases {
		_, err := createRedisConnection(&tc.redisConfig, log)
		assert.Equalf(t, tc.err, err, "Test[%d] FAILED, Could not connect to SQL, got error: %v\n", i, err)
	}
}

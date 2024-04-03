package database

import (
	"github.com/LetsFocus/goLF/logger"
	"github.com/LetsFocus/goLF/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_InitializeRedis(t *testing.T) {

	testcases := []struct {
		desc        string
		log         *logger.CustomLogger
		redisConfig *RedisConfig
		err         error
		output      Redis
	}{
		{
			desc: "Redis connected successfully",
			log:  logger.NewCustomLogger(),
			redisConfig: &RedisConfig{
				Host:           "127.0.0.1",
				Port:           "6379",
				Addr:           "127.0.0.1:6379",
				Password:       "",
				DB:             0,
				Retries:        5,
				RetryTime:      -1,
				PoolSize:       10,
				MinIdleConns:   4,
				MaxIdleConns:   8,
				ConnMaxLife:    time.Duration(10),
				ConMaxIdleTime: time.Duration(30),
			},
			err: nil,
		},
		{
			desc: "Redis Configure data not sufficient",
			log:  logger.NewCustomLogger(),
			redisConfig: &RedisConfig{
				Host:           "",
				Port:           "",
				Addr:           "",
				Password:       "",
				DB:             0,
				Retries:        5,
				RetryTime:      -1,
				PoolSize:       10,
				MinIdleConns:   4,
				MaxIdleConns:   8,
				ConnMaxLife:    time.Duration(10),
				ConMaxIdleTime: time.Duration(30),
			},
			err:    nil,
			output: Redis{},
		},
	}

	for i, tc := range testcases {
		_, err := InitializeRedis(tc.log, tc.redisConfig)
		assert.Equalf(t, tc.err, err, "Test[%d] FAILED, Could not connect to Redis, got error: %v\n", i, err)
	}
}

func Test_InitializeRedisFail(t *testing.T) {

	testcases := []struct {
		desc        string
		log         *logger.CustomLogger
		redisConfig *RedisConfig
		errString   string
		output      Redis
	}{

		{
			desc: "Redis failed to connect",
			log:  logger.NewCustomLogger(),
			redisConfig: &RedisConfig{
				Host:           "127.0.0.1",
				Port:           "6379",
				Addr:           "127.0.0.1:6377",
				Password:       "",
				DB:             0,
				Retries:        5,
				RetryTime:      -1,
				PoolSize:       10,
				MinIdleConns:   4,
				MaxIdleConns:   8,
				ConnMaxLife:    time.Duration(10),
				ConMaxIdleTime: time.Duration(30),
			},
			errString: "dial tcp 127.0.0.1:6377",
		},
	}

	for _, tc := range testcases {
		_, err := InitializeRedis(tc.log, tc.redisConfig)
		assert.Contains(t, err.Error(), tc.errString, "Testcase failed")
	}
}

func Test_RedisGetters(t *testing.T) {
	db := RedisConfig{
		Host:           "127.0.0.1",
		Port:           "6379",
		Addr:           "127.0.0.1:6379",
		Password:       "",
		DB:             0,
		Retries:        5,
		RetryTime:      -1,
		PoolSize:       10,
		MinIdleConns:   4,
		MaxIdleConns:   8,
		ConnMaxLife:    time.Duration(10),
		ConMaxIdleTime: time.Duration(30),
	}
	assert.Equal(t, db.GetHost(), db.Host)
	assert.Equal(t, db.GetMaxRetries(), db.Retries)
	assert.Equal(t, db.GetMaxRetryDuration(), db.RetryTime)
	assert.Equal(t, db.GetDBName(), RedisDB)
}

func Test_HealthCheckRedis(t *testing.T) {

	testcases := []struct {
		desc           string
		log            *logger.CustomLogger
		redisConfig    *RedisConfig
		expectedOutput types.Health
	}{
		{
			desc: "Redis connected successfully",
			log:  logger.NewCustomLogger(),
			redisConfig: &RedisConfig{
				Host:           "127.0.0.1",
				Port:           "6379",
				Addr:           "127.0.0.1:6379",
				Password:       "",
				DB:             0,
				Retries:        5,
				RetryTime:      -1,
				PoolSize:       10,
				MinIdleConns:   4,
				MaxIdleConns:   8,
				ConnMaxLife:    time.Duration(10),
				ConMaxIdleTime: time.Duration(30),
			},
			expectedOutput: types.Health{
				Status: Up,
				Name:   RedisDB,
			},
		},
		{
			desc: "Redis Configure data not sufficient",
			log:  logger.NewCustomLogger(),
			redisConfig: &RedisConfig{
				Host:           "",
				Port:           "",
				Addr:           "",
				Password:       "",
				DB:             0,
				Retries:        5,
				RetryTime:      -1,
				PoolSize:       10,
				MinIdleConns:   4,
				MaxIdleConns:   8,
				ConnMaxLife:    time.Duration(10),
				ConMaxIdleTime: time.Duration(30),
			},
			expectedOutput: types.Health{
				Status: Down,
				Name:   RedisDB,
			},
		},
		{
			desc: "Redis failed to connect",
			log:  logger.NewCustomLogger(),
			redisConfig: &RedisConfig{
				Host:           "127.0.0.1",
				Port:           "6379",
				Addr:           "127.0.0.1:6377",
				Password:       "",
				DB:             0,
				Retries:        5,
				RetryTime:      -1,
				PoolSize:       10,
				MinIdleConns:   4,
				MaxIdleConns:   8,
				ConnMaxLife:    time.Duration(10),
				ConMaxIdleTime: time.Duration(30),
			},
			expectedOutput: types.Health{
				Status: Down,
				Name:   RedisDB,
			},
		},
	}

	for i, tc := range testcases {
		redis, _ := InitializeRedis(tc.log, tc.redisConfig)
		assert.Equalf(t, tc.expectedOutput, redis.HealthCheckRedis(), "Test[%d] failed", i)
	}
}

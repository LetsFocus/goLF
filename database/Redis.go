package database

import (
	"context"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/LetsFocus/goLF/logger"
	"github.com/LetsFocus/goLF/types"
)

type RedisConfig struct {
	Host           string
	Port           string
	Addr           string
	Password       string
	DB             int
	Retries        int
	RetryTime      int
	PoolSize       int
	MinIdleConns   int
	MaxIdleConns   int
	ConMaxIdleTime time.Duration
	ConnMaxLife    time.Duration
}

func (r RedisConfig) GetHost() string {
	return r.Host
}

func (r RedisConfig) GetDBName() string {
	return RedisDB
}
func (r RedisConfig) GetMaxRetries() int {
	return r.Retries
}
func (r RedisConfig) GetMaxRetryDuration() int {
	return r.RetryTime
}

func InitializeRedis(log *logger.CustomLogger, c *RedisConfig) (Redis, error) {
	if c.Host != "" && c.Port != "" {
		client, err := createRedisConnection(c, log)
		if err != nil {
			return Redis{}, err
		}

		return Redis{Redis: client}, nil
	}

	return Redis{}, nil
}

func createRedisConnection(config *RedisConfig, log *logger.CustomLogger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:            config.Addr,
		Password:        config.Password,
		DB:              config.DB,
		MaxRetries:      config.Retries,
		MaxRetryBackoff: time.Duration(config.RetryTime) * time.Second,
		PoolSize:        config.PoolSize,
		MinIdleConns:    config.MinIdleConns,
		MaxIdleConns:    config.MaxIdleConns,
		ConnMaxIdleTime: config.ConMaxIdleTime * time.Minute,
		ConnMaxLifetime: config.ConnMaxLife * time.Second,
	})

	// PING to check if Redis is running
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Unable to connect the redis server %v", err)
		return nil, err
	}
	log.Info("Successfully connected to redis ")
	return client, nil
}

func (r Redis) HealthCheckRedis() types.Health {
	if isEmptyStruct(r) {
		return types.Health{Status: Down, Name: RedisDB}
	}

	if _, err := r.Redis.Ping(context.Background()).Result(); err != nil {
		return types.Health{Status: Down, Name: RedisDB}
	}

	return types.Health{Status: Up, Name: RedisDB}
}

func isEmptyStruct(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}

package database

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
)

type redisConfig struct {
	addr           string
	password       string
	dB             int
	retries        int
	retryTime      time.Duration
	poolSize       int
	minIdleConns   int
	maxIdleConns   int
	conMaxIdleTime time.Duration
	connMaxLife    time.Duration
}

func InitializeRedis(golf *model.GoLF, prefix string) {
	host := golf.Config.Get(prefix + "REDIS_HOST")
	port := golf.Config.Get(prefix + "REDIS_PORT")
	pwd := golf.Config.Get(prefix + "REDIS_PASSWORD")
	d := golf.Config.Get(prefix + "REDIS_DB_NUMBER")

	retry, err := strconv.Atoi(golf.Config.Get(prefix + "REDIS_MAX_RETRIES"))
	if err != nil {
		retry = 5
	}

	retryTime, err := time.ParseDuration(golf.Config.Get(prefix + "REDIS_RETRY_TIME"))
	if err != nil {
		retryTime = time.Duration(5)
	}

	poolSize, err := strconv.Atoi(golf.Config.Get(prefix + "REDIS_POOL_SIZE"))
	if err != nil {
		poolSize = 10
	}

	minIdleConns, err := strconv.Atoi(golf.Config.Get(prefix + "REDIS_MIN_IDLE_CONNS"))
	if err != nil {
		minIdleConns = 4
	}

	maxIdleConns, err := strconv.Atoi(golf.Config.Get(prefix + "REDIS_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = 8
	}

	conMaxIdleTime, err := time.ParseDuration(golf.Config.Get(prefix + "REDIS_CONN_MAX_IDLE_TIME"))
	if err != nil {
		conMaxIdleTime = time.Duration(30)
	}

	conMaxLife, err := time.ParseDuration(golf.Config.Get(prefix + "REDIS_CONN_MAX_LIFE"))
	if err != nil {
		conMaxLife = time.Duration(10)
	}
	db, err := strconv.Atoi(d)
	if err != nil && d != "" {
		golf.Logger.Errorf("invalid db number: %v", err)
		return
	}

	if host != "" && port != "" {
		redisCon := redisConfig{
			addr:           host + ":" + port,
			password:       pwd,
			dB:             db,
			retries:        retry,
			retryTime:      retryTime,
			poolSize:       poolSize,
			minIdleConns:   minIdleConns,
			maxIdleConns:   maxIdleConns,
			conMaxIdleTime: conMaxIdleTime,
			connMaxLife:    conMaxLife,
		}

		client, err := createRedisConnection(&redisCon, golf.Logger)
		if err == nil {
			golf.Redis = client
		}
	}
}

func createRedisConnection(config *redisConfig, log *logger.CustomLogger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:            config.addr,
		Password:        config.password,
		DB:              config.dB,
		MaxRetries:      config.retries,
		MaxRetryBackoff: config.retryTime * time.Second,
		PoolSize:        config.poolSize,
		MinIdleConns:    config.minIdleConns,
		MaxIdleConns:    config.maxIdleConns,
		ConnMaxIdleTime: config.conMaxIdleTime * time.Minute,
		ConnMaxLifetime: config.connMaxLife * time.Second,
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

package database

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
)

type redisConfig struct {
	addr     string
	password string
	dB       int
}

func InitializeRedis(golf *model.GoLF, prefix string) {
	localHost := golf.Config.Get(prefix + "REDIS_LOCALHOST")
	port := golf.Config.Get(prefix + "REDIS_PORT")
	pwd := golf.Config.Get(prefix + "REDIS_PASSWORD")
	d := golf.Config.Get(prefix + "REDIS_DB")

	db, err := strconv.Atoi(d)
	if err != nil && d != "" {
		golf.Logger.Errorf("invalid db number: %v", err)
		return
	}

	if localHost != "" && port != "" {
		redisCon := redisConfig{
			addr:     localHost + ":" + port,
			password: pwd,
			dB:       db,
		}

		client, err := createRedisConnection(&redisCon, golf.Logger)
		if err == nil {
			golf.Redis = client
		}
	}

}
func createRedisConnection(config *redisConfig, log *logger.CustomLogger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.addr,
		Password: config.password,
		DB:       config.dB,
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

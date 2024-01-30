package database

import (
	"context"
	"github.com/LetsFocus/goLF/logger"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/LetsFocus/goLF/goLF/model"
)

type redisConfig struct {
	addr     string
	password string
	dB       int
}

func InitializeRedis(golf *model.GoLF, prefix string) {
	db, _ := strconv.Atoi(golf.Config.Get(prefix + "REDIS_DB"))
	redisCon := redisConfig{
		addr:     golf.Config.Get(prefix + "REDIS_ADDR"),
		password: golf.Config.Get(prefix + "REDIS_PASSWORD"),
		dB:       db,
	}

	client, err := createRedisConnection(&redisCon, golf.Logger)
	if err == nil {
		golf.Redis = client
	}

}
func createRedisConnection(config *redisConfig, log *logger.CustomLogger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.addr,
		Password: config.password,
		DB:       config.dB,
	})

	// PING to check if Redis is running

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Unable to connect the redis server %v", err)
		return nil, err
	}
	log.Infof("Successfully connected to redis %v", pong)
	return client, nil
}

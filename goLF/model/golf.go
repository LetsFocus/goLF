package model

import (
	"database/sql"

	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/metrics"
	"github.com/LetsFocus/goLF/slogs"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
)

type GoLF struct {
	Database
	Config  configs.Config
	Logger  slogs.Log
	Metrics *metrics.Metrics
}
type RedisDB struct {
	Redis        *redis.Client
	RedisCLuster *redis.ClusterClient
}

type Database struct {
	Postgres *sql.DB
	RedisDB
	Elasticsearch *elasticsearch.Client
}

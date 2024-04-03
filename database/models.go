package database

import (
	"database/sql"
	"github.com/LetsFocus/goLF/types"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gocql/gocql"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	DB
	Redis
	Es
	Cassandra
}

type Redis struct {
	Redis        *redis.Client
	RedisCLuster *redis.ClusterClient
}

const (
	Up   = "up"
	Down = "down"
)

type DB struct {
	*sql.DB
	HealthCheck
}

type Es struct {
	*elasticsearch.Client
	HealthCheck
}

type Cassandra struct {
	*gocql.Session
	HealthCheck
}

type HealthCheck func() types.Health

type RetryFunc func(db *Database) error

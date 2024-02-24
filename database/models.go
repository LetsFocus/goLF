package database

import (
	"database/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gocql/gocql"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Redis        *redis.Client
	RedisCLuster *redis.ClusterClient
}

type Database struct {
	DB
	Redis
	Es
	CassandraClient
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

type CassandraClient struct {
	Session *gocql.Session
	Cluster *gocql.ClusterConfig
	HealthCheck
}

type HealthCheck func() Health

type Health struct {
	Host   string
	Status string
	Port   string
	Name   string
}

package model

import (
	"database/sql"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/logger"
	"github.com/redis/go-redis/v9"
)

type GoLF struct {
	Database
	Config configs.Config
	Logger *logger.CustomLogger
}

type Database struct {
	Postgres *sql.DB
	Redis    *redis.Client
}

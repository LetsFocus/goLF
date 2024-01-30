package model

import (
	"database/sql"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/logger"
)

type GoLF struct {
	Database
	Config configs.Config
	Logger *logger.CustomLogger
}

type Database struct {
	Postgres *sql.DB
}

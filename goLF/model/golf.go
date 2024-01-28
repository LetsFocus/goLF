package model

import (
	"database/sql"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/logger"
)

type GoLF struct {
	Conn   *sql.DB
	Config configs.Config
	Logger *logger.CustomLogger
}

package model

import (
	"database/sql"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/slogs"
)

type GoLF struct {
	Database
	Config configs.Config
	Logger slogs.Log
}

type Database struct {
	Postgres *sql.DB
}

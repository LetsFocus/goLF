package goLF

import (
	"database/sql"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/logger"
)

type GoLF struct {
	Conn   *sql.DB
	C      configs.Config
	Logger *logger.CustomLogger
}

func New() GoLF {
	var (
		goLF GoLF
		err  error
	)

	goLF.Logger = logger.NewCustomLogger()
	goLF.C = configs.NewConfig(goLF.Logger)

	goLF.Conn, err = database.InitializeDB(goLF.C, "")
	if err != nil {
		return goLF
	}

	return goLF
}

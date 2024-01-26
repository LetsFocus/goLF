package goLF

import (
	"strconv"
	"database/sql"
	"github.com/elastic/go-elasticsearch/v8"

	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/elasticstack"
	"github.com/LetsFocus/goLF/logger"
)

type GoLF struct {
	Conn   *sql.DB
	Elastic *elasticsearch.Client
	Config configs.Config
	Logger *logger.CustomLogger
}

func New() (GoLF, error) {
	var (
		goLF GoLF
		err  error
	)

	goLF.Logger = logger.NewCustomLogger()
	goLF.Config = configs.NewConfig(goLF.Logger)

	goLF.Conn, err = database.InitializeDB(goLF.Config, "")
	if err != nil {
		return goLF, err
	}

	retryCounter, err := strconv.Atoi(goLF.Config.Get(""+"ELASTIC_RETRY"))
	if err != nil {
		return goLF, err
	}

	goLF.Elastic, err = elasticstack.InitializeES(goLF.Config, "", retryCounter)
	if err != nil {
		return goLF, err
	}

	return goLF, nil
}

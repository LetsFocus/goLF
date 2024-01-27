package goLF

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/elasticstack"
	"github.com/LetsFocus/goLF/logger"
)

type Databases struct {
	Postgre *sql.DB
	Elastic *elasticsearch.Client
}

type GoLF struct {
	Conn    Databases
	Config  configs.Config
	Logger  *logger.CustomLogger
}

func New() GoLF {
	var goLF GoLF

	goLF.Logger = logger.NewCustomLogger()
	goLF.Config = configs.NewConfig(goLF.Logger)

	goLF.Conn.Postgre, _ = database.InitializeDB(goLF.Config, "")
	
	retry := goLF.Config.Get("" + "ELASTICSEARCH_RETRY")
	retryCounter, err := strconv.Atoi(retry)
	if err!=nil{
		retryCounter = 5
	}

	goLF.Conn.Elastic, _ = elasticstack.InitializeES(goLF.Config, "", retryCounter)
	return goLF
}

func Monitor(goLF *GoLF) {
	ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

	retry := goLF.Config.Get("" + "ELASTICSEARCH_RETRY")
	retryCounter, err := strconv.Atoi(retry)
	if err!=nil{
		retryCounter = 5
	}
	
	goLF.Config.Log.Info("Monitoring elasting search cluster...")
	for range ticker.C{
		goLF.Conn.Elastic, _ = elasticstack.MonitorES(goLF.Config, goLF.Conn.Elastic, "", retryCounter)
    }
}

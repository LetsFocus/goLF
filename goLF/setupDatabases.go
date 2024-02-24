package goLF

import (
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/goLF/model"
	"strconv"
)

func initializeDatabases(g *model.GoLF) {
	initializeSQL(g, "")
	initializeRedis()
	initializeCassandra()
	initializeElasticSearch()
}

func initializeSQL(g *model.GoLF, prefix string) {
	cfs := getSqlConfigs(g.Config, prefix)

	database.InitializeDB(cfs, g.Logger)
}

func getSqlConfigs(c configs.Config, prefix string) database.DBConfig {
	var (
		maxConnections, maxIdleConnections, connectionMaxLifeTime, idleConnectionTimeout, retry, retryTime int
		monitoring                                                                                         bool
		err                                                                                                error
	)

	maxIdleConnections, err = strconv.Atoi(c.Get(prefix + "DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		maxIdleConnections = 5
	}

	maxConnections, err = strconv.Atoi(c.Get(prefix + "DB_MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = 20
	}

	connectionMaxLifeTime, err = strconv.Atoi(c.Get(prefix + "DB_CONNECTIONS_MAX_LIFETIME"))
	if err != nil {
		connectionMaxLifeTime = 15
	}

	idleConnectionTimeout, err = strconv.Atoi(c.Get(prefix + "DB_IDLE_CONNECTIONS_TIMEOUT"))
	if err != nil {
		idleConnectionTimeout = 10
	}

	monitoring, err = strconv.ParseBool(c.Get(prefix + "DB_MONITORING"))
	if err != nil {
		monitoring = false
	}

	retry, err = strconv.Atoi(c.Get(prefix + "DB_RETRY_COUNT"))
	if err != nil {
		retry = 5
	}

	retryTime, err = strconv.Atoi(c.Get(prefix + "DB_RETRY_DURATION"))
	if err != nil {
		retryTime = 5
	}

	cfs := database.DBConfig{
		Host:                  c.Get(prefix + "DB_HOST"),
		Password:              c.Get(prefix + "DB_PASSWORD"),
		User:                  c.Get(prefix + "DB_USER"),
		Port:                  c.Get(prefix + "DB_PORT"),
		Dialect:               c.Get(prefix + "DB_DIALECT"),
		DBName:                c.Get(prefix + "DB_NAME"),
		SslMode:               c.Get(prefix + "DB_SSL"),
		MaxOpenConns:          maxConnections,
		MaxIdleConns:          maxIdleConnections,
		ConnMaxLifeTime:       connectionMaxLifeTime,
		IdleConnectionTimeout: idleConnectionTimeout,
		MonitoringEnable:      monitoring,
		Retry:                 retry,
		RetryDuration:         retryTime,
	}

	return cfs
}

func initializeElasticSearch() {

}

func initializeRedis() {

}

func initializeCassandra() {

}

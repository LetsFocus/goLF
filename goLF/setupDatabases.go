package goLF

import (
	"strconv"
	"strings"
	"time"

	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/goLF/model"
)

func initializeDatabases(g *model.GoLF) {
	InitializeSQL(g, "")
	InitializeRedis(g, "")
	InitializeCassandra(g, "")
	InitializeElasticSearch(g, "")
}

type DBConfig interface {
	GetMaxRetries() int
	GetMaxRetryDuration() int
	GetDBName() string
}

type RetryFunc func(g *model.GoLF, c DBConfig) (database.HealthCheck, error)

func initializeDatabase(g *model.GoLF, config DBConfig, retry RetryFunc) {
	if config == nil {
		g.Logger.Errorf("Configs are invalid")
		return
	}

	healthcheck, err := retry(g, config)
	if err != nil {
		g.Logger.Errorf("unable to establish %v connection with err:%v", config.GetDBName(), err)
		return
	}

	if healthcheck().Status == database.Up {
		g.Logger.Errorf("Successfully connected to %v", config.GetDBName())
	}

	go Monitoring(g, config, retry, healthcheck)
}

func InitializeSQL(g *model.GoLF, prefix string) {
	c := getSqlConfigs(g.Config, prefix)
	initializeDatabase(g, c, getSQLConnection)
}

func getSQLConnection(g *model.GoLF, c DBConfig) (healthcheck database.HealthCheck, err error) {
	dbConfig, _ := c.(database.DBConfig)

	db, err := database.InitializeDB(g.Logger, &dbConfig)
	if err != nil {
		return
	}

	g.Database.DB = db
	healthcheck = db.HealthCheckSQL
	return
}

func getESConnection(g *model.GoLF, c DBConfig) (healthcheck database.HealthCheck, err error) {
	dbConfig, _ := c.(database.ESConfig)

	es, err := database.InitializeES(g.Logger, &dbConfig)
	if err != nil {
		return
	}

	g.Database.Es = es
	healthcheck = es.HealthCheckES
	return
}

func getRedisConnection(g *model.GoLF, c DBConfig) (healthcheck database.HealthCheck, err error) {
	dbConfig, _ := c.(database.RedisConfig)

	redis, err := database.InitializeRedis(g.Logger, &dbConfig)
	if err != nil {
		return
	}

	g.Database.Redis = redis
	healthcheck = redis.HealthCheckRedis
	return
}

func getCassandraConnection(g *model.GoLF, c DBConfig) (healthcheck database.HealthCheck, err error) {
	dbConfig, _ := c.(database.CassandraConfig)

	redis, err := database.InitializeCassandra(g.Logger, &dbConfig)
	if err != nil {
		return
	}

	g.Database.Cassandra = redis
	healthcheck = redis.HealthCheckCassandra
	return
}

func Monitoring(g *model.GoLF, c DBConfig, retry RetryFunc, healthcheck database.HealthCheck) {
	ticker := time.NewTicker(time.Second)

	var (
		err          error
		retryCounter int
	)

monitoringLoop:
	for range ticker.C {
		if healthcheck().Status == database.Down {
			if retryCounter < c.GetMaxRetries() {
				for i := 0; i < c.GetMaxRetryDuration(); i++ {
					healthcheck, err = retry(g, c)
					if err == nil || healthcheck().Status == database.Up {
						retryCounter = 0
						break
					}

					retryCounter++
					time.Sleep(time.Second * time.Duration(c.GetMaxRetryDuration()))
					g.Logger.Errorf("%v Retry %d failed: %v", i+1, c.GetDBName(), err)
				}
			} else {
				break monitoringLoop
			}
		} else {
			retryCounter = 0
		}
	}

	ticker.Stop()
	g.Logger.Errorf("%v monitoring stopped after reaching maximum retries. Error for %v breakdown is %v", err)
}

func InitializeElasticSearch(g *model.GoLF, prefix string) {
	c := getESConfigs(g.Config, prefix)
	initializeDatabase(g, c, getESConnection)
}

func InitializeRedis(g *model.GoLF, prefix string) {
	c := getRedisConfigs(g.Config, prefix)
	initializeDatabase(g, c, getRedisConnection)
}

func InitializeCassandra(g *model.GoLF, prefix string) {
	c := getCassandraConfigs(g.Config, prefix)
	initializeDatabase(g, c, getRedisConnection)
}

func cleanAddresses(addressesString string) []string {
	addressesList := strings.Split(addressesString, ",")

	var addressesSlice []string

	for _, address := range addressesList {
		trimmedAddress := strings.TrimSpace(address)
		addressesSlice = append(addressesSlice, trimmedAddress)
	}

	return addressesSlice
}

func getRedisConfigs(c configs.Config, prefix string) database.RedisConfig {
	host := c.Get(prefix + "REDIS_HOST")
	port := c.Get(prefix + "REDIS_PORT")
	password := c.Get(prefix + "REDIS_PASSWORD")
	dbStr := c.Get(prefix + "REDIS_DB_NUMBER")

	retry, err := strconv.Atoi(c.Get(prefix + "REDIS_RETRY_COUNT"))
	if err != nil {
		retry = 5
	}

	retryTime, err := strconv.Atoi(c.Get(prefix + "REDIS_RETRY_DURATION"))
	if err != nil {
		retryTime = 5
	}

	poolSize, err := strconv.Atoi(c.Get(prefix + "REDIS_POOL_SIZE"))
	if err != nil {
		poolSize = 10
	}

	minIdleConns, err := strconv.Atoi(c.Get(prefix + "REDIS_MIN_IDLE_CONNECTIONS"))
	if err != nil {
		minIdleConns = 4
	}

	maxIdleConns, err := strconv.Atoi(c.Get(prefix + "REDIS_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		maxIdleConns = 8
	}

	conMaxIdleTime, err := time.ParseDuration(c.Get(prefix + "REDIS_IDLE_CONNECTIONS_TIMEOUT"))
	if err != nil {
		conMaxIdleTime = time.Duration(30)
	}

	conMaxLife, err := time.ParseDuration(c.Get(prefix + "REDIS_CONNECTIONS_MAX_LIFETIME"))
	if err != nil {
		conMaxLife = time.Duration(10)
	}

	db, _ := strconv.Atoi(dbStr)

	redisCfs := database.RedisConfig{
		Host:           host,
		Port:           port,
		Addr:           host + ":" + port,
		Password:       password,
		DB:             db,
		Retries:        retry,
		RetryTime:      retryTime,
		PoolSize:       poolSize,
		MinIdleConns:   minIdleConns,
		MaxIdleConns:   maxIdleConns,
		ConMaxIdleTime: conMaxIdleTime,
		ConnMaxLife:    conMaxLife,
	}

	return redisCfs
}

func getCassandraConfigs(c configs.Config, prefix string) database.CassandraConfig {
	var (
		maxConnections, connectionTimeout, maxRetries, retryDuration int
		monitoring                                                   bool
		err                                                          error
	)

	maxRetries, err = strconv.Atoi(c.Get(prefix + "CASSANDRA_RETRY_COUNT"))
	if err != nil {
		maxRetries = 5
	}

	retryDuration, err = strconv.Atoi(c.Get(prefix + "CASSANDRA_RETRY_DURATION"))
	if err != nil {
		retryDuration = 5
	}

	monitoring, err = strconv.ParseBool(c.Get(prefix + "CASSANDRA_MONITORING"))
	if err != nil {
		monitoring = false
	}

	maxConnections, err = strconv.Atoi(c.Get(prefix + "CASSANDRA_MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = 20
	}

	connectionTimeout, err = strconv.Atoi(c.Get(prefix + "CASSANDRA_TIMEOUT"))
	if err != nil {
		connectionTimeout = 100
	}

	addressesString := c.Get(prefix + "CASSANDRA_ADDRESSES")

	cfs := database.CassandraConfig{
		Addresses:         cleanAddresses(addressesString),
		Password:          c.Get(prefix + "CASSANDRA_PASSWORD"),
		MaxRetries:        maxRetries,
		DB:                c.Get(prefix + "CASSANDRA_DB"),
		RetryDuration:     retryDuration,
		MonitoringEnable:  monitoring,
		MaxOpenConns:      maxConnections,
		ConnectionTimeout: connectionTimeout,
	}

	return cfs
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

func getESConfigs(c configs.Config, prefix string) database.ESConfig {
	var (
		maxConnections, maxIdleConnections, idleConnectionTimeout, maxRetries, retryDuration int
		monitoring                                                                           bool
		err                                                                                  error
	)

	maxRetries, err = strconv.Atoi(c.Get(prefix + "ES_RETRY_COUNT"))
	if err != nil {
		maxRetries = 5
	}

	retryDuration, err = strconv.Atoi(c.Get(prefix + "ES_RETRY_DURATION"))
	if err != nil {
		retryDuration = 5
	}

	monitoring, err = strconv.ParseBool(c.Get(prefix + "ES_MONITORING"))
	if err != nil {
		monitoring = false
	}

	maxIdleConnections, err = strconv.Atoi(c.Get(prefix + "ES_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		maxIdleConnections = 5
	}

	maxConnections, err = strconv.Atoi(c.Get(prefix + "ES_MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = 20
	}

	idleConnectionTimeout, err = strconv.Atoi(c.Get(prefix + "ES_IDLE_CONNECTION_TIMEOUT"))
	if err != nil {
		idleConnectionTimeout = 10
	}

	addressesString := c.Get(prefix + "ES_ADDRESSES")

	cfs := database.ESConfig{
		Addresses:             cleanAddresses(addressesString),
		Username:              c.Get(prefix + "ES_USERNAME"),
		Password:              c.Get(prefix + "ES_PASSWORD"),
		MaxRetries:            maxRetries,
		RetryDuration:         retryDuration,
		MonitoringEnable:      monitoring,
		MaxOpenConns:          maxConnections,
		MaxIdleConns:          maxIdleConnections,
		IdleConnectionTimeout: idleConnectionTimeout,
	}

	return cfs
}

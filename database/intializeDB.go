package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/LetsFocus/goLF/errors"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
)

type dbConfig struct {
	host                  string
	password              string
	user                  string
	port                  string
	dialect               string
	dbName                string
	sslMode               string
	maxOpenConns          int
	maxIdleConns          int
	connMaxLifeTime       int
	idleConnectionTimeout int
	monitoringEnable      bool
	retry                 int
	retryDuration         int
}

func InitializeDB(golf *model.GoLF, prefix string) {
	var (
		maxConnections, maxIdleConnections, connectionMaxLifeTime, idleConnectionTimeout, retry, retryTime int
		monitoring                                                                                         bool
		err                                                                                                error
	)

	maxIdleConnections, err = strconv.Atoi(golf.Config.Get(prefix + "DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		maxIdleConnections = 5
	}

	maxConnections, err = strconv.Atoi(golf.Config.Get(prefix + "DB_MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = 20
	}

	connectionMaxLifeTime, err = strconv.Atoi(golf.Config.Get(prefix + "DB_CONNECTIONS_MAX_LIFETIME"))
	if err != nil {
		connectionMaxLifeTime = 15
	}

	idleConnectionTimeout, err = strconv.Atoi(golf.Config.Get(prefix + "DB_IDLE_CONNECTION_TIMEOUT"))
	if err != nil {
		idleConnectionTimeout = 10
	}

	monitoring, err = strconv.ParseBool(golf.Config.Get(prefix + "DB_MONITORING"))
	if err != nil {
		monitoring = false
	}

	retry, err = strconv.Atoi(golf.Config.Get(prefix + "DB_RETRY_COUNT"))
	if err != nil {
		retry = 5
	}

	retryTime, err = strconv.Atoi(golf.Config.Get(prefix + "DB_RETRY_DURATION"))
	if err != nil {
		retryTime = 5
	}

	c := dbConfig{
		host:                  golf.Config.Get(prefix + "DB_HOST"),
		password:              golf.Config.Get(prefix + "DB_PASSWORD"),
		user:                  golf.Config.Get(prefix + "DB_USER"),
		port:                  golf.Config.Get(prefix + "DB_PORT"),
		dialect:               golf.Config.Get(prefix + "DB_DIALECT"),
		dbName:                golf.Config.Get(prefix + "DB_NAME"),
		sslMode:               golf.Config.Get(prefix + "DB_SSL"),
		maxOpenConns:          maxConnections,
		maxIdleConns:          maxIdleConnections,
		connMaxLifeTime:       connectionMaxLifeTime,
		idleConnectionTimeout: idleConnectionTimeout,
		monitoringEnable:      monitoring,
	}

	if c.host != "" && c.port != "" && c.user != "" && c.password != "" && c.dialect != "" {
		if c.sslMode == "" {
			c.sslMode = "disable"
		}

		db, err := establishDBConnection(golf.Logger, c)
		if err == nil {
			db.SetMaxOpenConns(c.maxOpenConns)
			db.SetMaxIdleConns(c.maxIdleConns)
			db.SetConnMaxLifetime(time.Minute * time.Duration(c.connMaxLifeTime))
			db.SetConnMaxIdleTime(time.Minute * time.Duration(c.idleConnectionTimeout))

			golf.Postgres = db

			c.retry = retry
			c.retryDuration = retryTime
			if c.monitoringEnable {
				go monitoringDB(golf, c, c.retry, c.retryDuration)
			}
		}
	}
}

func monitoringDB(golf *model.GoLF, c dbConfig, retry, retryTime int) {
	ticker := time.NewTicker(time.Second)

	var (
		db           *sql.DB
		err          error
		retryCounter int
	)

monitoringLoop:
	for range ticker.C {
		if err = golf.Postgres.Ping(); err != nil {
			if retryCounter < retry {
				for i := 0; i < retry; i++ {
					db, err = establishDBConnection(golf.Logger, c)
					if err == nil {
						golf.Postgres = db
						retryCounter = 0

						break
					}

					retryCounter++
					time.Sleep(time.Second * time.Duration(retryTime))
					golf.Logger.Errorf("DB Retry %d failed: %v", i+1, err)
				}
			} else {
				break monitoringLoop
			}
		} else {
			retryCounter = 0
		}
	}

	ticker.Stop()
	golf.Logger.Errorf("DB Monitoring stopped after reaching maximum retries. Error for DB breakdown is %v", err)
}

func GenerateConnectionString(c dbConfig) string {
	switch c.dialect {
	case "mysql":
	case "postgres":
		return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", c.user, c.password, c.host, c.port, c.dbName, c.sslMode)
	}

	return ""
}

func establishDBConnection(log *logger.CustomLogger, c dbConfig) (*sql.DB, error) {
	connectionString := GenerateConnectionString(c)
	if connectionString == "" {
		log.Error("invalid dialect given")
		return nil, errors.Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError),
			Reason: "Invalid dialect"}
	}

	db, err := sql.Open(c.dialect, connectionString)
	if err != nil {
		log.Errorf("Failed to initialize the DB, Error:%v", err)
		return db, err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("Failed to ping the DB, Error:%v", err)
		return db, err
	}

	log.Info("database is connected successfully")
	return db, nil
}

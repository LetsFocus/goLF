package database

import (
	"database/sql"
	"fmt"
	"github.com/LetsFocus/goLF/errors"
	"github.com/LetsFocus/goLF/goLF/model"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type dbConfig struct {
	host             string
	password         string
	user             string
	port             string
	dialect          string
	dbName           string
	maxOpenConns     int
	maxIdleConns     int
	connMaxLifeTime  int
	monitoringEnable bool
	retry            int
	retryDuration    int
}

func InitializeDB(golf model.GoLF, prefix string) {
	var (
		maxConnections, maxIdleConnections, connectionMaxLifeTime, retry, retryTime int
		monitoring                                                                  bool
		err                                                                         error
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
		host:             golf.Config.Get(prefix + "DB_HOST"),
		password:         golf.Config.Get(prefix + "DB_PASSWORD"),
		user:             golf.Config.Get(prefix + "DB_USER"),
		port:             golf.Config.Get(prefix + "DB_PORT"),
		dialect:          golf.Config.Get(prefix + "DB_DIALECT"),
		dbName:           golf.Config.Get(prefix + "DB_NAME"),
		maxOpenConns:     maxConnections,
		maxIdleConns:     maxIdleConnections,
		connMaxLifeTime:  connectionMaxLifeTime,
		monitoringEnable: monitoring,
	}

	if c.host != "" && c.port != "" && c.user != "" && c.password != "" && c.dialect != "" {
		db, err := establishDBConnection(golf, c)
		if err == nil {
			db.SetMaxOpenConns(c.maxOpenConns)
			db.SetMaxIdleConns(c.maxIdleConns)
			db.SetConnMaxLifetime(time.Minute * time.Duration(c.maxOpenConns))

			golf.Conn = db

			go monitoringDB(golf, c, retry, retryTime)
		}
	}
}

func monitoringDB(golf model.GoLF, c dbConfig, retry, retryTime int) {
	ticker := time.NewTicker(time.Second)

	var (
		db           *sql.DB
		err          error
		retryCounter int
	)

	for {
		select {
		case <-ticker.C:
			if err = golf.Conn.Ping(); err != nil {
				db, err = establishDBConnection(golf, c)
				if err != nil {
					time.Sleep(time.Second * time.Duration(retryTime))
					golf.Logger.Error("DB Retrying failed")
				}

				golf.Conn = db
			}
		}
	}
}

func generateConnectionString(c dbConfig) string {
	switch c.dialect {
	case "mysql":
	case "postgres":
		return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.user, c.password, c.host, c.port, c.dbName)
	}

	return ""
}

func establishDBConnection(golf model.GoLF, c dbConfig) (*sql.DB, error) {
	connectionString := generateConnectionString(c)
	if connectionString == "" {
		golf.Logger.Error("invalid dialect given")
		return nil, errors.Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError),
			Reason: "Invalid dialect"}
	}

	db, err := sql.Open(c.dialect, connectionString)
	if err != nil {
		golf.Logger.Errorf("Failed to initialize the DB, Error:%v", err)
		return db, err
	}

	err = db.Ping()
	if err != nil {
		golf.Logger.Errorf("Failed to ping the DB, Error:%v", err)
		return db, err
	}

	golf.Logger.Info("database is connected successfully")
	return db, nil
}

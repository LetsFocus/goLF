package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/errors"
)

type config struct {
	host            string
	password        string
	user            string
	port            string
	dialect         string
	dbName          string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifeTime int
}

func InitializeDB(configs configs.Config, prefix string) (*sql.DB, error) {
	var (
		maxConnections, maxIdleConnections, connectionMaxLifeTime int
		err                                                       error
	)

	maxIdleConnections, err = strconv.Atoi(configs.Get(prefix + "DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		maxIdleConnections = 5
	}

	maxConnections, err = strconv.Atoi(configs.Get(prefix + "DB_MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = 20
	}

	connectionMaxLifeTime, err = strconv.Atoi(configs.Get(prefix + "DB_CONNECTIONS_MAX_LIFETIME"))
	if err != nil {
		connectionMaxLifeTime = 15
	}

	c := config{
		host:            configs.Get(prefix + "DB_HOST"),
		password:        configs.Get(prefix + "DB_PASSWORD"),
		user:            configs.Get(prefix + "DB_USER"),
		port:            configs.Get(prefix + "DB_PORT"),
		dialect:         configs.Get(prefix + "DB_DIALECT"),
		dbName:          configs.Get(prefix + "DB_NAME"),
		maxOpenConns:    maxConnections,
		maxIdleConns:    maxIdleConnections,
		connMaxLifeTime: connectionMaxLifeTime,
	}

	if c.host != "" && c.port != "" && c.user != "" && c.password != "" && c.dialect != "" {
		connectionString := generateConnectionString(c)
		if connectionString == "" {
			return nil, errors.Errors{StatusCode: 500, Code: "Invalid Dialect", Reason: "invalid dialect given"}
		}

		db, err := sql.Open(c.dialect, connectionString)
		if err != nil {
			configs.Log.Errorf("Failed to initialize the DB, Error:%v", err)
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			configs.Log.Errorf("Failed to initialize the DB, Error:%v", err)
			return nil, err
		}

		db.SetMaxOpenConns(c.maxOpenConns)
		db.SetMaxIdleConns(c.maxIdleConns)
		db.SetConnMaxLifetime(time.Minute * time.Duration(c.maxOpenConns))

		configs.Log.Info("database is connected successfully")
		return db, nil
	}

	return nil, nil
}

func generateConnectionString(c config) string {
	switch c.dialect {
	case "mysql":
	case "postgres":
		return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.user, c.password, c.host, c.port, c.dbName)
	}

	return ""
}

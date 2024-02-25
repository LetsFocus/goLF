package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/LetsFocus/goLF/errors"
	"github.com/LetsFocus/goLF/logger"
	"github.com/LetsFocus/goLF/types"
)

type DBConfig struct {
	Host                  string
	Password              string
	User                  string
	Port                  string
	Dialect               string
	DBName                string
	SslMode               string
	MaxOpenConns          int
	MaxIdleConns          int
	ConnMaxLifeTime       int
	IdleConnectionTimeout int
	MonitoringEnable      bool
	Retry                 int
	RetryDuration         int
}

func (d DBConfig) GetHost() string {
	return d.Host
}

func (d DBConfig) GetDBName() string {
	return SQL
}
func (d DBConfig) GetMaxRetries() int {
	return d.Retry
}
func (d DBConfig) GetMaxRetryDuration() int {
	return d.RetryDuration
}

func InitializeDB(log *logger.CustomLogger, c *DBConfig) (DB, error) {
	if c.Host != "" && c.Port != "" && c.User != "" && c.Password != "" && c.Dialect != "" {
		if c.SslMode == "" {
			c.SslMode = "disable"
		}

		db, err := EstablishDBConnection(log, c)
		if err == nil {
			return DB{}, err
		}

		db.SetMaxOpenConns(c.MaxOpenConns)
		db.SetMaxIdleConns(c.MaxIdleConns)
		db.SetConnMaxLifetime(time.Minute * time.Duration(c.ConnMaxLifeTime))
		db.SetConnMaxIdleTime(time.Minute * time.Duration(c.IdleConnectionTimeout))

		sqlDB := DB{DB: db}

		return sqlDB, nil
	}

	return DB{}, nil
}

func (d *DB) HealthCheckSQL() types.Health {
	if d == nil {
		return types.Health{Status: Down, Name: SQL}
	}

	if err := d.DB.Ping(); err != nil {
		return types.Health{Status: Down, Name: SQL}
	}

	return types.Health{Status: Up, Name: SQL}
}

func GenerateConnectionString(c DBConfig) string {
	switch c.Dialect {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	case "postgres":
		return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", c.User, c.Password, c.Host, c.Port, c.DBName, c.SslMode)
	}

	return ""
}

func EstablishDBConnection(log *logger.CustomLogger, c DBConfig) (*sql.DB, error) {
	connectionString := GenerateConnectionString(c)
	if connectionString == "" {
		log.Error("invalid dialect given")
		return nil, errors.Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError),
			Reason: "Invalid dialect"}
	}

	db, err := sql.Open(c.Dialect, connectionString)
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

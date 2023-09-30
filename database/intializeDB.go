package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/errors"
)

type config struct {
	host     string
	password string
	user     string
	port     string
	dialect  string
	dbName   string
}

func InitializeDB(configs configs.Config, prefix string) (*sql.DB, error) {
	c := config{host: configs.Get(prefix + "DB_HOST"), password: configs.Get(prefix + "DB_PASSWORD"),
		user: configs.Get(prefix + "DB_USER"), port: configs.Get(prefix + "DB_PORT"), dbName: configs.Get(prefix + "DB_NAME")}

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

	configs.Log.Info("database is connected successfully")
	return db, nil
}

func generateConnectionString(c config) string {
	switch c.dialect {
	case "mysql":
	case "postgres":
		return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.user, c.password, c.host, c.port, c.dbName)
	}

	return ""
}

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/LetsFocus/goLF/configs"
)

type dbConfig struct {
	host     string
	password string
	user     string
	port     string
	dbName   string
}

func InitializeDB(c configs.Config, prefix string) (*sql.DB, error) {
	configs := dbConfig{host: c.Get(prefix + "DB_HOST"), password: c.Get(prefix + "DB_PASSWORD"),
		user: c.Get(prefix + "DB_USER"), port: c.Get(prefix + "DB_PORT"), dbName: c.Get(prefix + "DB_NAME")}

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", configs.user, configs.password, configs.host, configs.port, configs.dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		c.Log.Errorf("Failed to initialize the DB, Error:%v", err)
		return nil, err
	}

	c.Log.Info("database is connected successfully")
	return db, nil
}

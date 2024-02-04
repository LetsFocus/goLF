package database

import (
	goErr "errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/errors"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
)

func Test_establishDBConnection(t *testing.T) {
	log := logger.NewCustomLogger()

	testcases := []struct {
		desc     string
		dbConfig dbConfig
		err      error
	}{
		{
			desc: "successfully established mysql db connection",
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
				dialect: "postgres", dbName: "testdb", sslMode: "disable"},
		},

		{
			desc:     "connectionString empty",
			dbConfig: dbConfig{dialect: "mysql"},
			err: errors.Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError),
				Reason: "Invalid dialect"},
		},

		{
			desc: "error while pinging the db",
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "root", password: "password",
				dialect: "postgres", dbName: "testdb", sslMode: "require"},
			err: goErr.New("pq: SSL is not enabled on the server"),
		},
	}

	for i, tc := range testcases {
		_, err := establishDBConnection(log, tc.dbConfig)

		assert.Equalf(t, tc.err, err, "Test[%d] FAILED, Could not connect to SQL, got error: %v\n", i, err)
	}
}

func Test_InitializeDB(t *testing.T) {
	t.Setenv("DB_DIALECT", "postgres")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_USER", "postgres")

	testcases := []struct {
		input *model.GoLF
	}{
		{
			input: &model.GoLF{
				Config: configs.Config{Log: logger.NewCustomLogger()},
				Logger: logger.NewCustomLogger(),
			},
		},
	}

	for _, tc := range testcases {
		InitializeDB(tc.input, "")
	}
}

func Test_MonitoringDB(t *testing.T) {
	db, _ := establishDBConnection(logger.NewCustomLogger(), dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
		dialect: "postgres", dbName: "testdb", sslMode: "disable"})

	db.Close()

	testcases := []struct {
		desc      string
		input     *model.GoLF
		dbConfig  dbConfig
		retry     int
		retryTime int
	}{
		{
			desc: "successfully monitored the db",
			input: &model.GoLF{
				Database: model.Database{Postgres: db},
				Config:   configs.Config{Log: logger.NewCustomLogger()},
				Logger:   logger.NewCustomLogger(),
			},
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
				dialect: "postgres", dbName: "testdb", sslMode: "disable"},
			retry:     3,
			retryTime: 1,
		},

		{
			desc: "retry is less than retryCount",
			input: &model.GoLF{
				Database: model.Database{Postgres: db},
				Config:   configs.Config{Log: logger.NewCustomLogger()},
				Logger:   logger.NewCustomLogger(),
			},
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
				dialect: "postgres", dbName: "testdb", sslMode: "disable"},
			retry:     0,
			retryTime: 1,
		},

		{
			desc: "DbConfigs are invalid",
			input: &model.GoLF{
				Database: model.Database{Postgres: db},
				Config:   configs.Config{Log: logger.NewCustomLogger()},
				Logger:   logger.NewCustomLogger(),
			},
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
				dialect: "mysql", dbName: "testdb", sslMode: "disable"},
			retry:     3,
			retryTime: 1,
		},
	}

	for _, tc := range testcases {
		go monitoringDB(tc.input, tc.dbConfig, tc.retry, tc.retryTime)

		time.Sleep(time.Second * 3)
	}
}

package database

import (
	goErr "errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/LetsFocus/goLF/errors"
	"github.com/LetsFocus/goLF/logger"
)

func Test_establishDBConnection(t *testing.T) {
	log := logger.NewCustomLogger()

	testcases := []struct {
		desc     string
		dbConfig DBConfig
		err      error
	}{
		{
			desc: "successfully established postgres db connection",
			dbConfig: DBConfig{Host: "localhost", Port: "5432", User: "postgres", Password: "password",
				Dialect: "postgres", DBName: "testdb", SslMode: "disable"},
		},

		{
			desc: "successfully established mysql db connection",
			dbConfig: DBConfig{Host: "localhost", Port: "3306", User: "mysql", Password: "password",
				Dialect: "mysql", DBName: "testdb", SslMode: "disabled"},
		},

		{
			desc:     "connectionString empty",
			dbConfig: DBConfig{Dialect: "redis"},
			err: errors.Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError),
				Reason: "Invalid dialect"},
		},

		{
			desc: "error while pinging the db",
			dbConfig: DBConfig{Host: "localhost", Port: "5432", User: "root", Password: "password",
				Dialect: "postgres", DBName: "testdb", SslMode: "require"},
			err: goErr.New("pq: SSL is not enabled on the server"),
		},
	}

	for i, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := EstablishDBConnection(log, &tc.dbConfig)

			assert.Equalf(t, tc.err, err, "Test[%d] FAILED, Could not connect to SQL, got error: %v\n", i, err)
		})
	}
}

/*
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
		log       string
	}{
		{
			desc: "successfully monitored the db",
			input: &model.GoLF{
				Database: model.Database{Pg: db},
				Config:   configs.Config{Log: logger.NewCustomLogger()},
				Logger:   logger.NewCustomLogger(),
			},
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
				dialect: "postgres", dbName: "testdb", sslMode: "disable"},
			retry:     3,
			retryTime: 1,
			log:       "database is connected successfully",
		},

		{
			desc: "retry is less than retryCount",
			input: &model.GoLF{
				Database: model.Database{Pg: db},
				Config:   configs.Config{Log: logger.NewCustomLogger()},
				Logger:   logger.NewCustomLogger(),
			},
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
				dialect: "postgres", dbName: "testdb", sslMode: "disable"},
			retry:     0,
			retryTime: 1,
			log:       "DB Monitoring stopped after reaching maximum retries. Error for DB breakdown is sql: database is closed",
		},

		{
			desc: "DbConfigs are invalid",
			input: &model.GoLF{
				Database: model.Database{Pg: db},
				Config:   configs.Config{Log: logger.NewCustomLogger()},
				Logger:   logger.NewCustomLogger(),
			},
			dbConfig: dbConfig{host: "localhost", port: "5432", user: "postgres", password: "password",
				dialect: "mysql", dbName: "testdb", sslMode: "disable"},
			retry:     3,
			retryTime: 1,
			log:       "invalid dialect given",
		},
	}

	for _, tc := range testcases {
		go monitoringDB(tc.input, tc.dbConfig, tc.retry, tc.retryTime)

		time.Sleep(time.Second * 3)
	}
}
*/

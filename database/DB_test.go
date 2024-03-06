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

func Test_establishDBConnectionFail(t *testing.T) {
	log := logger.NewCustomLogger()

	testcases := []struct {
		desc     string
		dbConfig DBConfig
		err      string
	}{

		{
			desc: "error in details",
			dbConfig: DBConfig{Host: "localhost", Port: "5434", User: "postgres", Password: "password",
				Dialect: "postgres", DBName: "testdb", SslMode: "disable"},
			err: "dial tcp ",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := EstablishDBConnection(log, &tc.dbConfig)

			assert.Contains(t, err.Error(), tc.err, "Testcase failed")
		})
	}
}

func Test_InitializeDB(t *testing.T) {
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
			desc: "empty values",
		},
		{
			desc: "successfully established postgres db connection",
			dbConfig: DBConfig{Host: "localhost", Port: "5432", User: "postgres", Password: "password",
				Dialect: "postgres", DBName: "testdb"},
		},

		{
			desc: "successfully established mysql db connection",
			dbConfig: DBConfig{Host: "localhost", Port: "3306", User: "mysql", Password: "password",
				Dialect: "mysql", DBName: "testdb", SslMode: "disabled"},
		},
	}

	for i, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := InitializeDB(log, &tc.dbConfig)

			assert.Equalf(t, tc.err, err, "Test[%d] FAILED, Could not connect to SQL, got error: %v\n", i, err)
		})
	}
}

func Test_Getters(t *testing.T) {
	db := DBConfig{
		Host:          "localhost",
		Port:          "5432",
		User:          "postgres",
		Password:      "password",
		Retry:         5,
		RetryDuration: 10,
	}
	assert.Equal(t, db.GetHost(), db.Host)
	assert.Equal(t, db.GetMaxRetries(), db.Retry)
	assert.Equal(t, db.GetMaxRetryDuration(), db.RetryDuration)
	assert.Equal(t, db.GetDBName(), "sql")
}

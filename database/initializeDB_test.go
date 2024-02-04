package database

import (
	"database/sql"
	"github.com/LetsFocus/goLF/slogs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_establishDBConnection(t *testing.T) {
	log := slogs.NewLogger()

	testcases := []struct {
		desc     string
		dbConfig dbConfig
		output   sql.DB
		err      error
	}{
		{
			desc:     "successfully established mysql db connection",
			dbConfig: dbConfig{dialect: "mysql", user: "root", password: "password", dbName: "test_db", port: "3306", host: "localhost"},
		},
	}

	for i, tc := range testcases {
		db, err := establishDBConnection(log, tc.dbConfig)

		assert.Equalf(t, tc.output, db, "Test[%d] failed", i)
		assert.Equalf(t, tc.err, err, "Test[%d] failed", i)

	}

}

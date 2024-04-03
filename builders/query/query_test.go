package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSelectQuery(t *testing.T) {
	testCases := []struct {
		Name           string
		Table          string
		Columns        []string
		Conditions     map[string]interface{}
		ExpectedQuery  string
		ExpectedValues []interface{}
	}{
		{
			Name:           "Test without columns and conditions",
			Table:          "users",
			Columns:        nil,
			Conditions:     nil,
			ExpectedQuery:  "SELECT * FROM users",
			ExpectedValues: nil,
		},
		{
			Name:           "Test with columns and conditions",
			Table:          "users",
			Columns:        []string{"name", "age"},
			Conditions:     map[string]interface{}{"id": 1, "status": "active"},
			ExpectedQuery:  "SELECT name, age FROM users WHERE id = ? AND status = ?",
			ExpectedValues: []interface{}{1, "active"},
		},
		{
			Name:           "Test with slice condition",
			Table:          "users",
			Columns:        nil,
			Conditions:     map[string]interface{}{"id": []interface{}{1, 2, 3}},
			ExpectedQuery:  "SELECT * FROM users WHERE id IN (?, ?, ?)",
			ExpectedValues: []interface{}{1, 2, 3},
		},
		{
			Name:           "Test with empty conditions",
			Table:          "users",
			Columns:        []string{"name"},
			Conditions:     nil,
			ExpectedQuery:  "SELECT name FROM users",
			ExpectedValues: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			query, values := BuildSelectQuery(tc.Table, tc.Columns, tc.Conditions)

			assert.Equal(t, tc.ExpectedQuery, query, "Unexpected query")
			assert.ElementsMatch(t, tc.ExpectedValues, values, "Unexpected values")
		})
	}
}

func TestBuildInsertQuery(t *testing.T) {
	testCases := []struct {
		Name           string
		Table          string
		Data           map[string]interface{}
		ExpectedQuery  string
		ExpectedValues []interface{}
	}{
		{
			Name:           "Test with single column",
			Table:          "users",
			Data:           map[string]interface{}{"name": "John"},
			ExpectedQuery:  "INSERT INTO users (name) VALUES (?)",
			ExpectedValues: []interface{}{"John"},
		},
		{
			Name:           "Test with multiple columns",
			Table:          "users",
			Data:           map[string]interface{}{"name": "John", "age": 30},
			ExpectedQuery:  "INSERT INTO users (name, age) VALUES (?, ?)",
			ExpectedValues: []interface{}{"John", 30},
		},
		{
			Name:           "Test with empty data",
			Table:          "users",
			Data:           map[string]interface{}{},
			ExpectedQuery:  "INSERT INTO users () VALUES ()",
			ExpectedValues: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			query, values := BuildInsertQuery(tc.Table, tc.Data)

			assert.Equal(t, tc.ExpectedQuery, query, "Unexpected query")
			assert.ElementsMatch(t, tc.ExpectedValues, values, "Unexpected values")
		})
	}
}

func TestBuildUpdateQuery(t *testing.T) {
	testCases := []struct {
		Name           string
		Table          string
		Data           map[string]interface{}
		Conditions     map[string]interface{}
		ExpectedQuery  string
		ExpectedValues []interface{}
	}{
		{
			Name:           "Test with single column update",
			Table:          "users",
			Data:           map[string]interface{}{"name": "John"},
			Conditions:     map[string]interface{}{"id": 1},
			ExpectedQuery:  "UPDATE users SET name = ? WHERE id = ?",
			ExpectedValues: []interface{}{"John", 1},
		},
		{
			Name:           "Test with multiple columns update",
			Table:          "users",
			Data:           map[string]interface{}{"name": "John", "age": 30},
			Conditions:     map[string]interface{}{"id": 1},
			ExpectedQuery:  "UPDATE users SET age = ?, name = ? WHERE id = ?",
			ExpectedValues: []interface{}{30, "John", 1},
		},
		{
			Name:           "Test with empty data",
			Table:          "users",
			Data:           map[string]interface{}{},
			Conditions:     map[string]interface{}{"id": 1},
			ExpectedQuery:  "UPDATE users SET  WHERE id = ?",
			ExpectedValues: []interface{}{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			query, values := BuildUpdateQuery(tc.Table, tc.Data, tc.Conditions)

			assert.Equal(t, tc.ExpectedQuery, query, "Unexpected query")
			assert.ElementsMatch(t, tc.ExpectedValues, values, "Unexpected values")
		})
	}
}

func TestBuildDeleteQuery(t *testing.T) {
	testCases := []struct {
		Name           string
		Table          string
		Conditions     map[string]interface{}
		ExpectedQuery  string
		ExpectedValues []interface{}
	}{
		{
			Name:           "Test with single condition",
			Table:          "users",
			Conditions:     map[string]interface{}{"id": 1},
			ExpectedQuery:  "DELETE FROM users WHERE id = ?",
			ExpectedValues: []interface{}{1},
		},
		{
			Name:           "Test with multiple conditions",
			Table:          "users",
			Conditions:     map[string]interface{}{"id": 1, "status": "inactive"},
			ExpectedQuery:  "DELETE FROM users WHERE id = ? AND status = ?",
			ExpectedValues: []interface{}{1, "inactive"},
		},
		{
			Name:           "Test with empty conditions",
			Table:          "users",
			Conditions:     nil,
			ExpectedQuery:  "DELETE FROM users",
			ExpectedValues: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			query, values := BuildDeleteQuery(tc.Table, tc.Conditions)

			assert.Equal(t, tc.ExpectedQuery, query, "Unexpected query")
			assert.ElementsMatch(t, tc.ExpectedValues, values, "Unexpected values")
		})
	}
}

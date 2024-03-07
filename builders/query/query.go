package query

import (
	"fmt"
	"sort"
	"strings"
)

// BuildSelectQuery builds a SELECT query based on the provided parameters
func BuildSelectQuery(table string, columns []string, conditions map[string]interface{}) (string, []interface{}) {
	query := "SELECT "

	// Add columns
	if len(columns) == 0 {
		query += "*"
	} else {
		query += strings.Join(columns, ", ")
	}

	// Add table
	query += " FROM " + table

	var values []interface{}

	// Add conditions
	if len(conditions) > 0 {
		query += " WHERE "
		conditionsList := make([]string, 0, len(conditions))

		// Sort keys for predictable order
		keys := make([]string, 0, len(conditions))
		for col := range conditions {
			keys = append(keys, col)
		}
		sort.Strings(keys)

		for _, col := range keys {
			val := conditions[col]
			switch v := val.(type) {
			case []interface{}:
				placeholders := make([]string, len(v))
				for i := range v {
					placeholders[i] = "?"
					values = append(values, v[i])
				}
				conditionsList = append(conditionsList, fmt.Sprintf("%s IN (%s)", col, strings.Join(placeholders, ", ")))
			default:
				conditionsList = append(conditionsList, fmt.Sprintf("%s = ?", col))
				values = append(values, val)
			}
		}
		query += strings.Join(conditionsList, " AND ")
	}

	return query, values
}

// BuildInsertQuery builds an INSERT query based on the provided parameters
func BuildInsertQuery(table string, data map[string]interface{}) (string, []interface{}) {
	query := "INSERT INTO " + table
	var columns []string
	var placeholders []string
	var values []interface{}

	for col, val := range data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	query += " (" + strings.Join(columns, ", ") + ") "
	query += "VALUES (" + strings.Join(placeholders, ", ") + ")"

	return query, values
}

// BuildUpdateQuery builds an UPDATE query based on the provided parameters
func BuildUpdateQuery(table string, data map[string]interface{}, conditions map[string]interface{}) (string, []interface{}) {
	query := "UPDATE " + table + " SET "
	var sets []string
	var values []interface{}

	// Sort keys for predictable order
	keys := make([]string, 0, len(data))
	for col := range data {
		keys = append(keys, col)
	}
	sort.Strings(keys)

	// Add sets
	for _, col := range keys {
		val := data[col]
		sets = append(sets, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	query += strings.Join(sets, ", ")

	// Add conditions
	if len(conditions) > 0 {
		query += " WHERE "
		conditionsList := make([]string, 0, len(conditions))

		// Sort keys for predictable order
		keys = keys[:0]
		for col := range conditions {
			keys = append(keys, col)
		}
		sort.Strings(keys)

		for _, col := range keys {
			val := conditions[col]
			conditionsList = append(conditionsList, fmt.Sprintf("%s = ?", col))
			values = append(values, val)
		}
		query += strings.Join(conditionsList, " AND ")
	}

	return query, values
}

// BuildDeleteQuery builds a DELETE query based on the provided parameters
func BuildDeleteQuery(table string, conditions map[string]interface{}) (string, []interface{}) {
	query := "DELETE FROM " + table

	var values []interface{}

	// Add conditions
	if len(conditions) > 0 {
		query += " WHERE "
		conditionsList := make([]string, 0, len(conditions))

		// Sort keys for predictable order
		keys := make([]string, 0, len(conditions))
		for col := range conditions {
			keys = append(keys, col)
		}
		sort.Strings(keys)

		for _, col := range keys {
			val := conditions[col]
			conditionsList = append(conditionsList, fmt.Sprintf("%s = ?", col))
			values = append(values, val)
		}
		query += strings.Join(conditionsList, " AND ")
	}

	return query, values
}

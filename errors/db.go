package errors

// DBError represents a 500 Internal Server Error indicating a database error.
type DBError struct{}

func (e *DBError) Error() string {
	return "Database error"
}

// RowsEffectedError represents a 500 Internal Server Error indicating an error in the number of rows affected.
type RowsEffectedError struct{}

func (e *RowsEffectedError) Error() string {
	return "No Rows Affected"
}

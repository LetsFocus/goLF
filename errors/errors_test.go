package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFunction(t *testing.T) {
	tests := []struct {
		name          string
		errorFunc     func(...string) error
		input         []string
		expectedError Errors
	}{
		{
			name: "InternalServerError",
			errorFunc: func(args ...string) error {
				return InternalServerError(errors.New("database error"))
			},
			input: []string{},
			expectedError: Errors{
				StatusCode: http.StatusInternalServerError,
				Code:       http.StatusText(http.StatusInternalServerError),
				Reason:     "db error",
			},
		},
		{
			name: "ServiceCallError",
			errorFunc: func(args ...string) error {
				return ServiceCallError(errors.New("service error"))
			},
			input: []string{},
			expectedError: Errors{
				StatusCode: http.StatusInternalServerError,
				Code:       http.StatusText(http.StatusInternalServerError),
				Reason:     "server failure",
			},
		},
		{
			name: "RowsAffectedError",
			errorFunc: func(args ...string) error {
				return RowsAffectedError(errors.New("database error"))
			},
			input: []string{},
			expectedError: Errors{
				StatusCode: http.StatusInternalServerError,
				Code:       http.StatusText(http.StatusInternalServerError),
				Reason:     "server is down",
			},
		},
		{
			name: "InvalidBody",
			errorFunc: func(args ...string) error {
				return InvalidBody()
			},
			input: []string{},
			expectedError: Errors{
				StatusCode: http.StatusBadRequest,
				Code:       http.StatusText(http.StatusBadRequest),
				Reason:     "invalid body",
			},
		},
		{
			name: "UnMarshalError",
			errorFunc: func(args ...string) error {
				return UnMarshalError()
			},
			input: []string{},
			expectedError: Errors{
				StatusCode: http.StatusBadRequest,
				Code:       http.StatusText(http.StatusBadRequest),
				Reason:     "incorrect data format",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errorFunc(tt.input...)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

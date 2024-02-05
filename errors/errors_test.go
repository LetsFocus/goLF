package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFunctions(t *testing.T) {
	tests := []struct {
		name          string
		errorFunc     func([]string) error
		input         []string
		expectedError Errors
	}{
		{
			name:      "InvalidParam",
			errorFunc: InvalidParam,
			input:     []string{"param1", "param2"},
			expectedError: Errors{
				StatusCode: http.StatusBadRequest,
				Code:       http.StatusText(http.StatusBadRequest),
				Reason:     "parameter param1,param2 is invalid",
			},
		},
		{
			name:      "MissingParam",
			errorFunc: MissingParam,
			input:     []string{"param1", "param2"},
			expectedError: Errors{
				StatusCode: http.StatusBadRequest,
				Code:       http.StatusText(http.StatusBadRequest),
				Reason:     "parameter param1,param2 is required",
			},
		},
		{
			name:      "MissingHeaders",
			errorFunc: MissingHeaders,
			input:     []string{"header1", "header2"},
			expectedError: Errors{
				StatusCode: http.StatusBadRequest,
				Code:       http.StatusText(http.StatusBadRequest),
				Reason:     "parameter header1,header2 is required in header",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errorFunc(tt.input)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

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
			name: "EntityNotFound",
			errorFunc: func(args ...string) error {
				return EntityNotFound("user", "123")
			},
			input: []string{},
			expectedError: Errors{
				StatusCode: http.StatusNotFound,
				Code:       http.StatusText(http.StatusNotFound),
				Reason:     "no user found for 123",
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

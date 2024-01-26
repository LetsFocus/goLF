package errors_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	custErr "github.com/LetsFocus/goLF/errors"
)

func TestInvalidParam(t *testing.T) {
	err := custErr.InvalidParam([]string{"param1", "param2"})
	assert.Equal(t, http.StatusBadRequest, err.(custErr.Errors).StatusCode)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.(custErr.Errors).Code)
	assert.Equal(t, "parameter param1,param2 is invalid", err.(custErr.Errors).Reason)
}

func TestMissingParam(t *testing.T) {
	err := custErr.MissingParam([]string{"param1", "param2"})
	assert.Equal(t, http.StatusBadRequest, err.(custErr.Errors).StatusCode)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.(custErr.Errors).Code)
	assert.Equal(t, "parameter param1,param2 is required", err.(custErr.Errors).Reason)
}

func TestMissingHeaders(t *testing.T) {
	err := custErr.MissingHeaders([]string{"param1", "param2"})
	assert.Equal(t, http.StatusBadRequest, err.(custErr.Errors).StatusCode)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.(custErr.Errors).Code)
	assert.Equal(t, "parameter param1,param2 is required in header", err.(custErr.Errors).Reason)
}

func TestInternalServerError(t *testing.T) {
	err := custErr.InternalServerError(errors.New("database error"))
	assert.Equal(t, http.StatusInternalServerError, err.(custErr.Errors).StatusCode)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), err.(custErr.Errors).Code)
	assert.Equal(t, "db error", err.(custErr.Errors).Reason)
}

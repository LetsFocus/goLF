package errors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingHeader_Error(t *testing.T) {
	header := "exampleHeader"
	err := &MissingHeader{Header: header}
	expected := fmt.Sprintf("Missing header: '%s'", header)
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestMissingHeaders_Error(t *testing.T) {
	headers := []string{"header1", "header2", "header3"}
	err := &MissingHeaders{Headers: headers}
	expected := fmt.Sprintf("Missing headers: '%s'", strings.Join(headers, "', '"))
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestMissingParam_Error(t *testing.T) {
	param := "exampleParam"
	err := &MissingParam{Param: param}
	expected := fmt.Sprintf("Missing parameter: '%s'", param)
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestMissingParams_Error(t *testing.T) {
	params := []string{"param1", "param2", "param3"}
	err := &MissingParams{Params: params}
	expected := fmt.Sprintf("Missing parameters: '%s'", strings.Join(params, "', '"))
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

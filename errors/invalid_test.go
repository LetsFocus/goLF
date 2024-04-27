package errors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidParam_Error(t *testing.T) {
	param := "exampleParam"
	err := &InvalidParam{Param: param}
	expected := fmt.Sprintf("Invalid parameter: '%s'", param)
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestInvalidParams_Error(t *testing.T) {
	params := []string{"param1", "param2", "param3"}
	err := &InvalidParams{Params: params}
	expected := fmt.Sprintf("Invalid parameters: '%s'", strings.Join(params, "', '"))
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestInvalidHeader_Error(t *testing.T) {
	header := "exampleHeader"
	err := &InvalidHeader{Header: header}
	expected := fmt.Sprintf("Invalid header: '%s'", header)
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestInvalidHeaders_Error(t *testing.T) {
	headers := []string{"header1", "header2", "header3"}
	err := &InvalidHeaders{Headers: headers}
	expected := fmt.Sprintf("Invalid headers: '%s'", strings.Join(headers, "', '"))
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

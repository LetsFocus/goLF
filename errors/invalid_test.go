package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestInvalidParam_Error(t *testing.T) {
	param := "exampleParam"
	err := &InvalidParam{Param: param}
	expected := fmt.Sprintf("Invalid parameter: '%s'", param)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestInvalidParams_Error(t *testing.T) {
	params := []string{"param1", "param2", "param3"}
	err := &InvalidParams{Params: params}
	expected := fmt.Sprintf("Invalid parameters: '%s'", strings.Join(params, "', '"))
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestInvalidHeader_Error(t *testing.T) {
	header := "exampleHeader"
	err := &InvalidHeader{Header: header}
	expected := fmt.Sprintf("Invalid header: '%s'", header)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestInvalidHeaders_Error(t *testing.T) {
	headers := []string{"header1", "header2", "header3"}
	err := &InvalidHeaders{Headers: headers}
	expected := fmt.Sprintf("Invalid headers: '%s'", strings.Join(headers, "', '"))
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

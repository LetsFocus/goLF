package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestMissingHeader_Error(t *testing.T) {
	header := "exampleHeader"
	err := &MissingHeader{Header: header}
	expected := fmt.Sprintf("Missing header: '%s'", header)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestMissingHeaders_Error(t *testing.T) {
	headers := []string{"header1", "header2", "header3"}
	err := &MissingHeaders{Headers: headers}
	expected := fmt.Sprintf("Missing headers: '%s'", strings.Join(headers, "', '"))
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestMissingParam_Error(t *testing.T) {
	param := "exampleParam"
	err := &MissingParam{Param: param}
	expected := fmt.Sprintf("Missing parameter: '%s'", param)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestMissingParams_Error(t *testing.T) {
	params := []string{"param1", "param2", "param3"}
	err := &MissingParams{Params: params}
	expected := fmt.Sprintf("Missing parameters: '%s'", strings.Join(params, "', '"))
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

package errors

import (
	"fmt"
	"github.com/LetsFocus/goLF/logger"
	"net/http"
	"strings"
)

type Errors struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Code       string `json:"code,omitempty"`
	Reason     string `json:"reason,omitempty"`
}

func (e Errors) Error() string {
	return fmt.Sprintf("{\n code: %v \n reason: %v\n}", e.Code, e.Reason)
}

func InvalidParam(err []string) error {
	reason := fmt.Sprintf("parameter "+strings.Join(err, ",")+" is invalid", err)

	logger.NewCustomLogger().Errorf("Invalid Parameter: %v", strings.Join(err, ","))
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: reason}
}

func MissingParam(err []string) error {
	reason := fmt.Sprintf("parameter " + strings.Join(err, ",") + " is required")

	logger.NewCustomLogger().Errorf("Missing Parameter: %v", strings.Join(err, ","))
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: reason}
}

func MissingHeaders(err []string) error {
	reason := fmt.Sprintf("parameter " + strings.Join(err, ",") + " is required in header")

	logger.NewCustomLogger().Errorf("Missing Headers: %v", strings.Join(err, ","))
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: reason}
}

func InternalServerError(err error) error {

	logger.NewCustomLogger().Errorf("Error from DB, Error: %v", err)
	return Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError), Reason: "db error"}
}

func ServiceCallError(err error) error {

	logger.NewCustomLogger().Errorf("Error from the service, Error: %v", err)
	return Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError), Reason: "server failure"}
}

func RowsAffectedError(err error) error {
	logger.NewCustomLogger().Errorf("Error from DB, Error: %v", err)
	return Errors{StatusCode: http.StatusInternalServerError, Code: http.StatusText(http.StatusInternalServerError), Reason: "server is down"}
}

func InvalidBody() error {
	logger.NewCustomLogger().Errorf("Invalid Body")
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: "invalid body"}
}

func EntityNotFound(entity, id string) error {
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: fmt.Sprintf("no %v found for %v", entity, id)}
}

func UnMarshalError(entity, id string) error {
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: fmt.Sprintf("incorrect data format")}
}

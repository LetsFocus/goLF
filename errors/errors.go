package errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LetsFocus/goLF/logger"
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
	reason := fmt.Sprintf("parameter " + strings.Join(err, ",") + " is invalid")

	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: reason}
}

func MissingParam(err []string) error {
	reason := fmt.Sprintf("parameter " + strings.Join(err, ",") + " is required")

	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: reason}
}

func MissingHeaders(err []string) error {
	reason := fmt.Sprintf("parameter " + strings.Join(err, ",") + " is required in header")

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
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: "invalid body"}
}

func EntityNotFound(entity, id string) error {
	return Errors{StatusCode: http.StatusNotFound, Code: http.StatusText(http.StatusNotFound), Reason: fmt.Sprintf("no %v found for %v", entity, id)}
}

func UnMarshalError() error {
	return Errors{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: fmt.Sprintf("incorrect data format")}
}

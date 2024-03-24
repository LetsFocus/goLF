package errors

import (
	"fmt"
	"github.com/LetsFocus/goLF/logger"
	"net/http"
	"strings"
)

type HTTPError struct {
    StatusCode int    `json:"statusCode,omitempty"`
    Code       string `json:"code,omitempty"`
    Reason     string `json:"reason,omitempty"`
}

func (e HTTPError) Error() string {
    return fmt.Sprintf("{\n code: %v \n reason: %v\n}", e.Code, e.Reason)
}

func NewHTTPError(statusCode int, reason string) *HTTPError {
    return &HTTPError{
        StatusCode: statusCode,
        Code:       http.StatusText(statusCode),
        Reason:     reason,
    }
}

var (
    ErrInvalidParam = NewHTTPError(http.StatusBadRequest, "invalid parameter")
    ErrMissingParam = NewHTTPError(http.StatusBadRequest, "missing parameter")
    ErrMissingHeaders = NewHTTPError(http.StatusBadRequest, "missing header parameter")
    ErrInternalServerError = NewHTTPError(http.StatusInternalServerError, "internal server error")
    ErrServiceCallError = NewHTTPError(http.StatusInternalServerError, "service call error")
    ErrRowsAffectedError = NewHTTPError(http.StatusInternalServerError, "rows affected error")
    ErrInvalidBody = NewHTTPError(http.StatusBadRequest, "invalid body")
)

func NewEntityNotFoundError(entity, id string) *HTTPError {
    return &HTTPError{
        StatusCode: http.StatusNotFound,
        Code:       http.StatusText(http.StatusNotFound),
        Reason:     fmt.Sprintf("no %v found for %v", entity, id),
    }
}

func NewUnmarshalError() *HTTPError {
    return &HTTPError{
        StatusCode: http.StatusBadRequest,
        Code:       http.StatusText(http.StatusBadRequest),
        Reason:     "incorrect data format",
    }
}

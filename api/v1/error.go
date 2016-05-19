package v1

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	StatusCode() int
	error
}

type SimpleHTTPError struct {
	code    int
	message string
}

func (e *SimpleHTTPError) Error() string {
	return e.message
}
func (e *SimpleHTTPError) StatusCode() int {
	return e.code
}

func RouteNotFound(url string) HTTPError {
	return &SimpleHTTPError{code: http.StatusNotFound, message: fmt.Sprintf("No such route: %s", url)}
}
func BadRequest(reason string, args ...interface{}) HTTPError {
	return &SimpleHTTPError{code: http.StatusBadRequest, message: fmt.Sprintf("Bad Request: %s", fmt.Sprintf(reason, args...))}
}
func ServerError(reason string, args ...interface{}) HTTPError {
	return &SimpleHTTPError{code: http.StatusInternalServerError, message: fmt.Sprintf("Server error: %s", fmt.Sprintf(reason, args...))}
}

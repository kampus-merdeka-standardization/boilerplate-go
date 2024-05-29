package pkg_error

import (
	"fmt"
	"net/http"
)

type ClientError struct {
	Code    int
	Message string
	Raw     error
}

func (e ClientError) Error() string {
	return fmt.Sprintf("%d\t%s", e.Code, e.Message)
}

func NewBadRequest(err error, msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusBadRequest,
		Raw:     err,
		Message: msg,
	}
}

func NewNotFound(err error, msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusNotFound,
		Raw:     err,
		Message: msg,
	}
}

func NewForbidden(err error, msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusForbidden,
		Raw:     err,
		Message: msg,
	}
}

func NewUnauthorized(err error, msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusUnauthorized,
		Raw:     err,
		Message: msg,
	}
}

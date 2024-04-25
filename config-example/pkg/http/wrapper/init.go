package pkg_http_wrapper

import (
	"time"
)

func NewResponseWithValue[T any](message string, value T) response[T] {
	timestamp := time.Now().Format(time.RFC3339)

	return response[T]{
		Message:   message,
		Value:     value,
		Timestamp: timestamp,
	}
}

func NewResponse(message string) response[any] {
	timestamp := time.Now().Format(time.RFC3339)

	return response[any]{
		Message:   message,
		Timestamp: timestamp,
	}
}

func NewError(error string) response[any] {
	timestamp := time.Now().Format(time.RFC3339)
	return response[any]{
		Error:     error,
		Timestamp: timestamp,
	}
}

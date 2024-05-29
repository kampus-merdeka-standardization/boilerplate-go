package pkg_http_wrapper

import "time"

func NewResponseWithValue[T any](code int,message string, value T) response[T] {
	timestamp := time.Now().Format(time.RFC3339)

	return response[T]{
		Meta: meta{
			Code: code,
			Message: message,
			Timestamp: timestamp,
		},
		Value:     value,
	}
}

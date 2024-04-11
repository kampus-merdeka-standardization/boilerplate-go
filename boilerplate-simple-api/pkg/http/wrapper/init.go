package pkg_http_wrapper

import "time"

func NewResponse(message string, value any) response {
	timestamp := time.Now().Format(time.RFC3339)

	return response{
		Message:   message,
		Value:     value,
		Timestamp: timestamp,
	}
}

func NewError(error string) response {
	return response{
		Error: error,
	}
}

package pkg_http_wrapper

import "time"

func NewResponse(code int,message string) response[any] {
	timestamp := time.Now().Format(time.RFC3339)

	return response[any]{
		Meta: meta{
			Code: code,
			Message: message,
			Timestamp: timestamp,
		},
	}
}

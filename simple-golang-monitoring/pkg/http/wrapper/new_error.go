package pkg_http_wrapper

import "time"

func NewError(code int,err string) response[any] {
	timestamp := time.Now().Format(time.RFC3339)
	return response[any]{
		Meta: meta{
			Code: code,
			Message: err,
			Timestamp: timestamp,
		},	
	}
}
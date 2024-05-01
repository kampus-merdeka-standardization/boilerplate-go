package pkg_http_wrapper

type (
	response[T any] struct {
		Meta meta `json:"metadata"`
		Value     T      `json:"value,omitempty"`
	}
	meta struct {
		Message string `json:"message"`
		Code int `json:"code"`
		Timestamp string `json:"timestamp"`
	}

)

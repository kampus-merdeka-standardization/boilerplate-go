package pkg_http_wrapper

type (
	response[T any] struct {
		Message   string `json:"message,omitempty"`
		Value     T      `json:"value,omitempty"`
		Error     string `json:"error,omitempty"`
		Timestamp string `json:"timestamp"`
	}
)

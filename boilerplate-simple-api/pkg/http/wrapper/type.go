package pkg_http_wrapper

type (
	response struct {
		Message   string      `json:"message"`
		Value     interface{} `json:"value"`
		Error     string      `json:"error"`
		Timestamp string      `json:"timestamp"`
	}
)

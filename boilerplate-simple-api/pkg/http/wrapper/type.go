package pkg_http_wrapper

type (
	Response struct {
		Message string      `json:"message"`
		Value   interface{} `json:"value"`
	}

	Error struct {
		Message string `json:"message"`
	}
)

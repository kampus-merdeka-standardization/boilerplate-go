package pkg_http

type (
	Response struct {
		Message string      `json:"message"`
		Value   interface{} `json:"value,omitempty"`
	}

	Error struct {
		Message string `json:"message"`
	}

	TraceType string
)

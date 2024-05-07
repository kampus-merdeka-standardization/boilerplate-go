package mock_http

type Response struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type PostRequest struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

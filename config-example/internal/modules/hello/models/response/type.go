package hello_response

type GetHelloByName struct {
	Name  string `json:"name"`
	Hello string `json:"hello"`
}

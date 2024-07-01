package hello_request

type CreateHello struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=0"`
}

type ReplaceHello struct {
	CurrentName string `json:"current_name" binding:"required"`
	NewName     string `json:"new_name" binding:"required"`
}

type UpdateHello struct {
	NewName string `json:"new_name" binding:"required"`
}

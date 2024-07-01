package product_request

type UpdateProduct struct {
	Name  string `json:"name" binding:"required"`
	Price int64  `json:"price" binding:"required,gte=0"`
}

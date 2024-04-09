package product_request

type CreateProduct struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gte=0"`
}

type UpdateProduct struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gte=0"`
}

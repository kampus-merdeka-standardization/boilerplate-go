package product_request

type UpdateProduct struct {
	ID    string  `json:"id" binding:"required,uuid"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gte=0"`
}

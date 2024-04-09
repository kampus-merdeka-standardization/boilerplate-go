package product_api

import product_usecase "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/usecase"

type productController struct {
	productUsecase product_usecase.ProductUsecase
}

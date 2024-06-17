package product_api

import product_usecase "simple-golang-database/internal/modules/product/usecase"

type productController struct {
	productUsecase product_usecase.ProductUsecase
}

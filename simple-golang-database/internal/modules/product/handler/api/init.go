package product_api

import (
	product_usecase "simple-golang-database/internal/modules/product/usecase"

	"github.com/gin-gonic/gin"
)

func NewProductController(router *gin.RouterGroup, productUsecase product_usecase.ProductUsecase) {
	pController := &productController{
		productUsecase: productUsecase,
	}

	router.POST("", pController.CreateProduct)
	router.GET("/:id", pController.GetProductByID)
	router.GET("", pController.GetAllProduct)
	router.PUT("/:id", pController.UpdateProductByID)
	router.DELETE("/:id", pController.DeleteProductByID)
}

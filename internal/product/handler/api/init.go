package product_api

import (
	"github.com/gin-gonic/gin"
	product_usecase "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/usecase"
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

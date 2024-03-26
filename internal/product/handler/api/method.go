package product_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product_request "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/request"
	product_response "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/model/response"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
)

func (p *productController) GetAllProduct(ctx *gin.Context) {
	products, err := p.productUsecase.GetAllProduct(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	var productsResponse []product_response.Product
	for _, p := range products {
		productsResponse = append(productsResponse, product_response.Product{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		})
	}

	ctx.JSON(http.StatusOK, httpPkg.Response{
		Message: "Sucessfully Retrieved All Product",
		Value:   productsResponse,
	})
}

func (p *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := p.productUsecase.GetProductByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpPkg.Response{
		Message: "Successfully Retrieved Product By the ID of " + id,
		Value: product_response.Product{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		},
	})
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var req product_request.CreateProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	id, err := p.productUsecase.CreateProduct(ctx, req.Name, req.Price)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, httpPkg.Response{
		Message: "Successfully Create Product",
		Value: gin.H{
			"product_id": id,
		},
	})
}

func (p *productController) UpdateProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var req product_request.UpdateProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	err := p.productUsecase.UpdateProductByID(ctx, id, req.Name, req.Price)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpPkg.Response{
		Message: "Successfully Updated Product",
	})
}

func (p *productController) DeleteProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := p.productUsecase.DeleteProductByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpPkg.Response{
		Message: "Successfully Deleted Product",
	})
}

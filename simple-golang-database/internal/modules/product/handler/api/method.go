package product_api

import (
	"net/http"

	product_request "simple-golang-database/internal/modules/product/model/request"
	product_response "simple-golang-database/internal/modules/product/model/response"
	httpWrapper "simple-golang-database/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
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

	res := httpWrapper.NewResponseWithValue(http.StatusOK, "Successully Get All Product", productsResponse)

	ctx.JSON(http.StatusOK, res)
}

func (p *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := p.productUsecase.GetProductByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := httpWrapper.NewResponseWithValue(http.StatusOK, "Successully Get Product", product)

	ctx.JSON(http.StatusOK, res)
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

	res := httpWrapper.NewResponseWithValue(http.StatusOK, "Successfully Created Product", gin.H{
		"product_id": id,
	})

	ctx.JSON(http.StatusOK, res)
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

	res := httpWrapper.NewResponse(http.StatusOK, "Succesffuly Updated Product")

	ctx.JSON(http.StatusOK, res)
}

func (p *productController) DeleteProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := p.productUsecase.DeleteProductByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := httpWrapper.NewResponse(http.StatusOK, "Successfully Deleted Product")

	ctx.JSON(http.StatusOK, res)
}

package product_handler

import (
	"net/http"
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

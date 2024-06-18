package product_handler

import (
	"net/http"
	product_request "simple-golang-database/internal/modules/product/model/request"
	httpWrapper "simple-golang-database/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

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

package product_handler

import (
	"net/http"
	product_request "simple-golang-database/internal/modules/product/model/request"
	httpWrapper "simple-golang-database/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

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

	res := httpWrapper.NewResponse(http.StatusNoContent, "Succesffuly Updated Product")

	ctx.JSON(http.StatusNoContent, res)
}

package product_handler

import (
	"net/http"
	httpWrapper "simple-golang-database/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

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

package product_handler

import (
	"net/http"

	httpWrapper "simple-golang-database/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

func (p *productController) DeleteProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := p.productUsecase.DeleteProductByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := httpWrapper.NewResponse(http.StatusNoContent, "Successfully Deleted Product")

	ctx.JSON(http.StatusOK, res)
}

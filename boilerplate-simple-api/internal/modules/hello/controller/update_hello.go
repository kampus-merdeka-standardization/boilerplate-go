package hello_controller

import (
	"fmt"
	"net/http"
	hello_request "simple-api/internal/modules/hello/models/request"

	"github.com/gin-gonic/gin"
)

// path : /hello [PATCH]
func (hc *helloController) UpdateHello(ctx *gin.Context) {
	var req hello_request.UpdateHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		fmt.Sprintf("Your name is replaced to %s", req.NewName),
	)
}

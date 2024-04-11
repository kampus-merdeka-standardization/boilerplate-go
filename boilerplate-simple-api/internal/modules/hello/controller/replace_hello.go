package hello_controller

import (
	"fmt"
	"net/http"
	hello_request "simple-api/internal/modules/hello/models/request"

	"github.com/gin-gonic/gin"
)

// path : /hello [PUT]
func (hc *helloController) ReplaceHello(ctx *gin.Context) {
	var req hello_request.ReplaceHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		fmt.Sprintf("Your name is replaced from %s to %s", req.CurrentName, req.NewName),
	)
}

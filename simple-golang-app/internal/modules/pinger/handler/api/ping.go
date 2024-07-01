package pinger_api

import (
	httpPkg "simple-golang-app/pkg/http"

	"github.com/gin-gonic/gin"
)

func (p *pingerController) Ping(ctx *gin.Context) {
	ctx.JSON(200, httpPkg.Response{
		Message: "pong",
	})
}

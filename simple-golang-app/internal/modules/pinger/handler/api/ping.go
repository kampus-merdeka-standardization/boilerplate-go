package pinger_api

import (
	"github.com/gin-gonic/gin"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/http"
)

func (p *pingerController) Ping(ctx *gin.Context) {
	ctx.JSON(200, httpPkg.Response{
		Message: "pong",
	})
}

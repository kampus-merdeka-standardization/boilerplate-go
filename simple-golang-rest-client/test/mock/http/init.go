package mock_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	srv := gin.Default()

	gin.EnableJsonDecoderDisallowUnknownFields()

	product := srv.Group("/product")

	product.GET(":id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			ID:    ctx.Param("id"),
			Name:  "Sikat Gigi",
			Price: 12500,
		})
	})

	product.POST("", func(ctx *gin.Context) {
		var req PostRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{
				Message: "Unable to Parse JSON",
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(200, Response{
			ID:    "002",
			Name:  req.Name,
			Price: req.Price,
		})
	})

	srv.Run(":8080")
}

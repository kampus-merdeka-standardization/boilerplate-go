package hello_controller

import (
	"fmt"
	"net/http"
	hello_request "simple-api/internal/modules/hello/models/request"
	"time"

	"github.com/gin-gonic/gin"
)

// path : /hello/:name [GET]
func (hc *helloController) GetHelloByName(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("Hello, %s!", ctx.Param("name")),
	})
}

// path : /hello [POST]
func (hc *helloController) CreateHello(ctx *gin.Context) {
	var req hello_request.CreateHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message":   fmt.Sprintf("Hello, %s you are %d years old", req.Name, req.Age),
			"timestamp": time.Now().Format(time.RFC3339),
		},
	})
}

// path : /hello [PUT]
func (hc *helloController) ReplaceHello(ctx *gin.Context) {
	var req hello_request.ReplaceHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message":   fmt.Sprintf("Your name is replaced from %s to %s", req.CurrentName, req.NewName),
			"timestamp": time.Now().Format(time.RFC3339),
		},
	})
}

// path : /hello [PATCH]
func (hc *helloController) UpdateHello(ctx *gin.Context) {
	var req hello_request.UpdateHello
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message":   fmt.Sprintf("Your name is replaced to %s", req.NewName),
			"timestamp": time.Now().Format(time.RFC3339),
		},
	})
}

// path : /hello/:id [DELETE]
func (hc *helloController) DeleteHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   fmt.Sprintf("Your data by the id of %s is deleted", ctx.Param("id")),
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

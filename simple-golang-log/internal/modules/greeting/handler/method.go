package greeting_handler

import (
	"errors"
	"fmt"
	pkg_http "simple-golang-log/pkg/http"

	"github.com/gin-gonic/gin"
)

func (gh *greetingHandler) GreetByName(ctx *gin.Context) {
	name := ctx.Param("name")

	if name == "erma" {
		ctx.Error(errors.New("no greeting for mbak erma >:)"))
		return
	}

	ctx.JSON(200, pkg_http.Response{
		Message: fmt.Sprintf("Hello %s!", name),
	})
}

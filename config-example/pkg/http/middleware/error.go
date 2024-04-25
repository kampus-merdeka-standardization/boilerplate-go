package pkg_http_middleware

import (
	"github.com/gin-gonic/gin"

	errorPkg "config-example/pkg/error"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0]
		// if err can be casted to ClientError, then it is a client error
		if clientError, ok := err.Err.(*errorPkg.ClientError); ok {
			c.JSON(clientError.Code, pkg_http_wrapper.NewError(
				clientError.Message,
			))
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			c.JSON(400, pkg_http_wrapper.NewError(err.Error()))
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			c.JSON(500, pkg_http_wrapper.NewError("Internal Server Error"))
			return
		}

		c.JSON(500, pkg_http_wrapper.NewError("Internal Server Error"))
	}
}

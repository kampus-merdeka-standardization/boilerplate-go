package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	errorPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/error"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0]
		// if err can be casted to ClientError, then it is a client error
		if clientError, ok := err.Err.(*errorPkg.ClientError); ok {
			c.JSON(clientError.Code, httpPkg.Error{
				Message: clientError.Message,
			})
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			c.JSON(400, httpPkg.Error{
				Message: err.Err.Error(),
			})
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			fmt.Println(err.Err.Error())
			c.JSON(500, httpPkg.Error{
				Message: "Internal server error",
			})
			return
		}

		c.JSON(500, httpPkg.Error{
			Message: "Internal server error",
		})
		logger.Error(err.Error(), zap.Any("Error", err))
	}
}

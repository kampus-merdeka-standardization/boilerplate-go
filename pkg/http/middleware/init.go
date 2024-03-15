package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	errorPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/error"
	httpPkg "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]
			// if err can be casted to ClientError, then it is a client error
			if clientError, ok := err.Err.(*errorPkg.ClientError); ok {
				c.JSON(clientError.Code, httpPkg.Error{
					Message: clientError.Message,
				})
			} else if err.IsType(gin.ErrorTypeBind) {
				c.JSON(400, httpPkg.Error{
					Message: err.Err.Error(),
				})
			} else if err.IsType(gin.ErrorTypePrivate) {
				fmt.Println(err.Err.Error())
				c.JSON(500, httpPkg.Error{
					Message: "Internal server error",
				})
			} else {
				fmt.Println(err.Err.Error())
				c.JSON(500, httpPkg.Error{
					Message: "Internal server error",
				})
			}
		}
	}
}

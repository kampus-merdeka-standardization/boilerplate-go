package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	errorPkg "simple-golang-app/pkg/error"
	httpPkg "simple-golang-app/pkg/http"
	"simple-golang-app/pkg/logger"
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
			logger.Error(clientError.Raw.Error(), zap.Int("Code", clientError.Code))
			c.JSON(clientError.Code, httpPkg.Error{
				Message: clientError.Message,
			})
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			logger.Error(err.Error(), zap.Any("Error", err))
			c.JSON(400, httpPkg.Error{
				Message: err.Err.Error(),
			})
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			logger.Error(err.Error(), zap.Any("Error", err))
			c.JSON(500, httpPkg.Error{
				Message: "Internal server error",
			})
			return
		}

		logger.Error(err.Error(), zap.Any("Error", err))
		c.JSON(500, httpPkg.Error{
			Message: "Internal server error",
		})
	}
}

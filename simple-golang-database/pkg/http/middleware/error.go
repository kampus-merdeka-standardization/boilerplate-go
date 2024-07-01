package pkg_http_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	errorPkg "simple-golang-database/pkg/error"
	pkg_http_wrapper "simple-golang-database/pkg/http/wrapper"
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
			c.JSON(clientError.Code, pkg_http_wrapper.NewResponse(clientError.Code, clientError.Message))
			return
		}

		if validationErr, ok := err.Err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, pkg_http_wrapper.NewResponse(http.StatusBadRequest, validationErr.Error()))
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			c.JSON(400, pkg_http_wrapper.NewResponse(http.StatusBadRequest, err.Err.Error()))
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			c.JSON(500, pkg_http_wrapper.NewResponse(http.StatusInternalServerError, "Internal Server Error"))
			return
		}

		c.JSON(500, pkg_http_wrapper.NewResponse(http.StatusInternalServerError, "Internal Server Error"))
	}
}

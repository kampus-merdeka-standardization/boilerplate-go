// Package http handles the http webserver
package http

import (
	"io"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// NewHTTPServer returns gin http server
func NewHTTPServer(ginMode string) *gin.Engine {
	if ginMode == gin.ReleaseMode {
		gin.SetMode(ginMode)
	} else if ginMode == gin.TestMode {
		gin.SetMode(ginMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	if ve, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ve.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
	}

	logFile, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.New()

	router.Use()

	return router
}

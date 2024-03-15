package pinger_api

import "github.com/gin-gonic/gin"

func NewPingerController(router *gin.RouterGroup) *pingerController {
	pController := &pingerController{}

	router.GET("/ping", pController.Ping)

	return pController
}

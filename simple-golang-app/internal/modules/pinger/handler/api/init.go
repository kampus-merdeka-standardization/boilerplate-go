package pinger_api

import "github.com/gin-gonic/gin"

func NewPingerController(router *gin.RouterGroup) {
	pController := &pingerController{}

	router.GET("", pController.Ping)
}

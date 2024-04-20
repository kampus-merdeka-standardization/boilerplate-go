package hello_handler

import "github.com/gin-gonic/gin"

func BindHelloHandler(router *gin.RouterGroup) {
	// root path : /hello
	controller := new(helloController)

	router.GET("/:name", controller.GetHelloByName)
	router.POST("", controller.CreateHello)
	router.PUT("", controller.ReplaceHello)
	router.PATCH("", controller.UpdateHello)
	router.DELETE("/:id", controller.DeleteHello)
}

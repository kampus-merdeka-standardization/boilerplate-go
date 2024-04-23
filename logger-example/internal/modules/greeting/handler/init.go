package greeting_handler

import "github.com/gin-gonic/gin"

func BindGreetingHandler(router *gin.RouterGroup) {
	gHandler := new(greetingHandler)

	router.GET("/:name", gHandler.GreetByName)
}

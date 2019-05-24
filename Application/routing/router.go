package routing

import (
	"github.com/gin-gonic/gin"
)

func GenRouter() *gin.Engine {
	router := gin.Default()
	//v1 := router.Group("/v1")
	{
		//v1.GET("/hello_world/:id", actions.HelloWorldHandler)
	}
	return router
}

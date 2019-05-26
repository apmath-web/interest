package routing

import (
	"github.com/apmath-web/interests/Application/actions"
	"github.com/gin-gonic/gin"
)

func GenRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", actions.GetInerest)
	}
	return router
}

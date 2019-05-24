package actions

import (
	"github.com/apmath-web/interests/Application/viewModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInerest(c *gin.Context) {
	evm := viewModels.GenIdsViewModel(1, []int{})

	c.JSON(http.StatusOK, evm)
}

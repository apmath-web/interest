package actions

import (
	"net/http"
	"strconv"

	"github.com/apmath-web/interests/Application/viewModels"
	"github.com/apmath-web/interests/Infrastructure/repositories"
	"github.com/gin-gonic/gin"
)

func HelloWorldHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, viewModels.GenHelloWorldViewModel("Wrong id", "400"))
		return
	}
	hw := repositories.Repo.GetModel(id)
	if hw == nil {
		c.JSON(http.StatusNotFound, viewModels.GenHelloWorldViewModel("No model", "404"))
		return
	}
	helloWorldViewModel := viewModels.GenHelloWorldViewModel(hw.GetMessage(), "200")
	c.JSON(http.StatusOK, helloWorldViewModel)
}

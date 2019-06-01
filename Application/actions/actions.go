package actions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Application/viewModels"
	"github.com/apmath-web/interests/Domain/models"
	"github.com/apmath-web/interests/Infrastructure"
	"github.com/gin-gonic/gin"
)

func GetInterests(c *gin.Context) {

	clientId, err := strconv.Atoi(c.Param("clientId"))
	if err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage("client id must be numeric")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	if clientId < 0 {
		validator := Validation.GenValidation()
		validator.SetMessage("param error")
		validator.AddMessage(Validation.GenMessage("clientId", "Is negative"))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	vm := viewModels.IdsViewModel{}

	if err := c.BindJSON(&vm); err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage("body error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	if !vm.Validate() {
		validator := vm.GetValidation()
		validator.SetMessage("validation error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	for _, id := range vm.CoborrowersIdSlice {
		if clientId == id {
			validator := Validation.GenValidation()
			validator.SetMessage("validation error")
			validator.AddMessage(Validation.GenMessage("coBorrowers", "Client's ID is equal to coborrower's ID"))
			str, _ := json.Marshal(validator)
			c.String(http.StatusBadRequest, string(str))
			return
		}
	}

	dm := models.GenIds(clientId, vm.GetCoborrowersIdSlice())

	service := Infrastructure.GetServiceManager().GetCalculationService()

	ei, err := service.Calculate(dm)

	if err != nil {
		if err.Error() == "clients service not available" {
			c.String(http.StatusInternalServerError, string(err.Error()))
			return
		}
		if err.Error() == "bad request" {
			c.String(http.StatusBadRequest, string(err.Error()))
			return
		}
		if err.Error() == "client not found" {
			c.String(http.StatusNotFound, string(err.Error()))
			return
		}
		c.String(http.StatusBadRequest, string(err.Error()))
		return
	}

	evm := new(viewModels.InterestsViewModel)
	evm.Hydrate(ei)

	c.JSON(http.StatusOK, evm)
}

package mapper

import (
	"github.com/apmath-web/interests/Application/viewModels"
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Domain/models"
)

func IdsViewMapper(vm viewModels.IdsViewModel, clientId int) Domain.IdsDomainModelInterface {
	return models.GenIds(clientId, vm.GetCoborrowersIdSlice())
}

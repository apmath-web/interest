package Mappers

import (
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Domain/models"
	"github.com/apmath-web/interests/Infrastructure/applicationModels"
)

func PersonApplicationMapper(am applicationModels.PersonApplicationModel) Domain.PersonDomainModelInterface {
	DomainModel := new(models.PersonDomainModel)
	DomainModel.FirstName = am.FirstName
	DomainModel.LastName = am.LastName
	DomainModel.SumWage = am.GetSumWage()
	return DomainModel
}

package mappers

import (
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Domain/models"
	"github.com/apmath-web/interests/Infrastructure/applicationModels"
)

func PersonDomainModelMapper(personApplicationModel applicationModels.PersonApplicationModel) Domain.PersonDomainModelInterface {
	personDomainModel := new(models.PersonDomainModel)
	personDomainModel.FirstName = personApplicationModel.FirstName
	personDomainModel.LastName = personApplicationModel.LastName
	personDomainModel.SumWage = personApplicationModel.GetSumWage()
	return personDomainModel
}

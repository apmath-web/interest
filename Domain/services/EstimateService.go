package services

import (
	"errors"
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Domain/models"
)

type EstimateService struct{}

func (es *EstimateService) Calculate(persons []Domain.PersonDomainModelInterface) (Domain.InterestsInterface, error) {

	sumWagePerson := 0

	for _, value := range persons {
		if value.GetSumWage() <= 0 {
			return nil, errors.New("Wage is negative value or zero")
		}
		sumWagePerson += value.GetSumWage()
	}

	dm := models.GenInterestsDomainModel(5.0)

	switch {
	case sumWagePerson >= 500000:
		dm.SetInterest(10.0)
	case sumWagePerson >= 100000:
		dm.SetInterest(8.0)
	case sumWagePerson >= 50000:
		dm.SetInterest(6.0)
	case sumWagePerson >= 30000:
		dm.SetInterest(4.0)
	}

	return dm, nil
}

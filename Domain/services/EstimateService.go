package services

import (
	"errors"
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Domain/models"
)

type CalculateService struct {
}

func (cs *CalculateService) Estimate(persons []Domain.PersonDomainModelInterface) (Domain.InterestsInterface, error) {
	var sumWagePerson int

	for _, value := range persons {
		if value.GetSumWage() <= 0 {
			return nil, errors.New("Wage is negative value or zero")
		}
		sumWagePerson += value.GetSumWage()
	}

	interest := models.GenInterestsDomainModel(7.5)

	switch {
	case sumWagePerson >= 500000:
		interest.SetInterest(12.5)
	case sumWagePerson >= 100000:
		interest.SetInterest(10.0)
	case sumWagePerson >= 50000:
		interest.SetInterest(8.5)
	}
	return interest, nil
}

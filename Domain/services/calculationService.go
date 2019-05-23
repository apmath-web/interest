package services

import (
	"github.com/apmath-web/interests/Domain"
	"github.com/apmath-web/interests/Domain/models"
)

type CalculationService struct {
	clientFetcher Domain.ClientFetchInterface
}

func Calculate(persons []Domain.PersonDomainModelInterface) float64 {
	return 7.5
} 

func (cs *CalculationService) Calculate(ids Domain.IdsDomainModelInterface) (Domain.InterestsInterface, error) {
	var persons []Domain.PersonDomainModelInterface

	var pdm, err = cs.clientFetcher.Fetch(ids.GetClientId())
	if err != nil {
		return nil, err
	}
	persons = append(persons, pdm)

	for _, value := range ids.GetCoborrowersIdSlice() {
		var pdm, err = cs.clientFetcher.Fetch(value)
		if err != nil {
			return nil, err
		}
		persons = append(persons, pdm)
	}

	var percent = Calculate(persons)
	
	var interests = models.GenInterestsDomainModel(percent)

	return interests, nil
}

func (cs *CalculationService) GenCalculationService(clientFetch Domain.ClientFetchInterface) {
	cs.clientFetcher = clientFetch
}


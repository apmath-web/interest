package services

import (
	"github.com/apmath-web/interests/Domain"
)

type CalculationService struct {
	clientFetcher    Domain.ClientFetchInterface
	calculateService CalculateService
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

	interest, err := cs.calculateService.Estimate(persons)

	if err != nil {
		return nil, err
	}

	return interest, nil
}

func GenCalculationService(clientFetch Domain.ClientFetchInterface, calc CalculateService) CalculationService {
	cs := new(CalculationService)
	cs.clientFetcher = clientFetch
	cs.calculateService = calc
	return *cs
}

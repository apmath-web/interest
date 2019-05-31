package services

import (
	"github.com/apmath-web/interests/Domain"
)

type CalculationService struct {
	clientFetcher Domain.ClientFetchInterface
}

func (cs *CalculationService) Calculate(ids Domain.IdsDomainModelInterface) (Domain.InterestsInterface, error) {
	var persons []Domain.PersonDomainModelInterface
	var calc = CalculateService{}
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

	interest, err := calc.Estimate(persons)

	if err != nil {
		return nil, err
	}

	return interest, nil
}

func (cs *CalculationService) GenCalculationService(clientFetch Domain.ClientFetchInterface) {
	cs.clientFetcher = clientFetch
}

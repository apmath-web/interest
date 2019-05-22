package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Domain"
)

type JsonInterests struct {
	amount float64 `json:"Interest"`
}

type InterestsViewModel struct {
	JsonInterests
	validation Validation.Validation
}

func (interestsViewModel *InterestsViewModel) GetInterest() float64 {
	return interestsViewModel.amount
}

func (interestsViewModel *InterestsViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"Interest": interestsViewModel.amount,
	})
}

func (interestsViewModel *InterestsViewModel) Hydrate(interestsInerface Domain.InterestsInterface) {
	interestsViewModel.amount = interestsInerface.GetInterest()
}

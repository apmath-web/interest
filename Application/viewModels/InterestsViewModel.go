package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Domain"
)

type JsonInterests struct {
	interest float64 `json:"interest"`
}

type InterestsViewModel struct {
	JsonInterests
	validation Validation.Validation
}

func (interestsViewModel *InterestsViewModel) GetInterest() float64 {
	return interestsViewModel.interest
}

func (interestsViewModel *InterestsViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"interest": interestsViewModel.interest,
	})
}

func (interestsViewModel *InterestsViewModel) Hydrate(interestsInerface Domain.InterestsInterface) {
	interestsViewModel.interest = interestsInterface.GetInterest()
}

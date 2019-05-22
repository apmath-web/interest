package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Domain"
)

type JsonInterests struct {
	amount float64 `json:"maxPayment"`
}

type InterestsViewModel struct {
	JsonInterests
	validation Validation.Validation
}

func (interestsViewModel *InterestsViewModel) GetAmount() float64 {
	return interestsViewModel.amount
}

func (interestsViewModel *InterestsViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"maxPayment": interestsViewModel.amount,
	})
}

func (interestsViewModel *InterestsViewModel) Hydrate(interestsInerface Domain.InterestsInterface) {
	interestsViewModel.amount = interestsInerface.GetAmount()
}

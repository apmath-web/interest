package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Domain"
)

type IdsViewModel struct {
	CoborrowersIdSlice []int `json:"coBorrowers"`
	validation Validation.Validation
}


func (idsViewModel *IdsViewModel) GetCoborrowersIdSlice() []int {
	return idsViewModel.CoborrowersIdSlice
}


func (idsViewModel *IdsViewModel) validateCoBorrowerIdSlice() bool {
	for _, id := range idsViewModel.CoborrowersIdSlice {
		if id < 0 {
			idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", "Is negative"))
			return false
		}
	}
	return true
}

func (idsViewModel *IdsViewModel) Validate() bool {
	if idsViewModel.validateCoBorrowerIdSlice() {
		return true
	}
	return false
}


func (idsViewModel *IdsViewModel) UnmarshalJSON(b []byte) error {
	tmpIds := idsViewModel.CoborrowersIdSlice
	err := json.Unmarshal(b, &tmpIds)
	if err := json.Unmarshal(b, &tmpIds); err != nil {
		return err
	}
	idsViewModel.CoborrowersIdSlice = tmpIds
	return err
}

func (idsViewModel *IdsViewModel) GetValidation() Domain.ValidationInterface {
	return &idsViewModel.validation
}

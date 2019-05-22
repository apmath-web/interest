package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Domain"
)

type JsonIds struct {
	ClientId           int   `json:"clientId"`    //ID заемщика
	CoborrowersIdSlice []int `json:"coBorrowers"` //Cписок ID созаемщиков
}

type IdsViewModel struct {
	JsonIds
	validation Validation.Validation
}

func (idsViewModel *IdsViewModel) GetClienId() int {
	return idsViewModel.ClientId
}

func (idsViewModel *IdsViewModel) GetCoborrowersIdSlice() []int {
	return idsViewModel.CoborrowersIdSlice
}

func (idsViewModel *IdsViewModel) validateClientId() bool {
	if idsViewModel.ClientId < 0 {
		idsViewModel.validation.AddMessage(Validation.GenMessage("clientId", "Is negative"))
		return false
	}
	return true
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
	if idsViewModel.validateClientId() && idsViewModel.validateCoBorrowerIdSlice() {
		return true
	}
	return false
}

func (idsViewModel *IdsViewModel) UnmarshalJSON(b []byte) error {
	tmpIds := JsonIds{}
	err := json.Unmarshal(b, &tmpIds)
	if err := json.Unmarshal(b, &tmpIds); err != nil {
		return err
	}
	idsViewModel.JsonIds = tmpIds
	return err
}

func (idsViewModel *IdsViewModel) GetValidation() Domain.ValidationInterface {
	return &idsViewModel.validation
}

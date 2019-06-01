package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Domain"
	"strconv"
)

type JsonIds struct {
	CoborrowersIdSlice []int `json:"coBorrowers"` //Cписок ID созаемщиков
}

type IdsViewModel struct {
	JsonIds
	validation Validation.Validation
}

func (idsViewModel *IdsViewModel) GetCoborrowersIdSlice() []int {
	return idsViewModel.CoborrowersIdSlice
}

func (idsViewModel *IdsViewModel) validateCoBorrowerIdSlice() bool {
	for _, id := range idsViewModel.CoborrowersIdSlice {
		if id < 0 {
			str := "ID is negative: " + strconv.Itoa(id)
			idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", str))
			return false
		}
	}
	for _, id := range idsViewModel.CoborrowersIdSlice {
		if id < 0 {
			str := "Is negative: " + strconv.Itoa(id)
			idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", str))
			return false
		}
	}
		if len(Validation.Unique(idsViewModel.CoborrowersIdSlice)) != len(idsViewModel.CoborrowersIdSlice) {
			idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", "IDs are equal to each other"))
			return false
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

package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/interests/Application/Validation"
	"github.com/apmath-web/interests/Domain"
)

type IdsViewModel struct {
	ClientId           int   `json:"clientId"`
	CoborrowersIdSlice []int `json:"coBorrowers"`
	validation         Validation.Validation
}

func (idsvm *IdsViewModel) SetClientId(clientId int) {
	idsvm.ClientId = clientId
}

func (idsvm *IdsViewModel) SetCoBorrowersIdSlice(coborrowersIdSlice []int) {
	idsvm.CoborrowersIdSlice = coborrowersIdSlice
}

func GenIdsViewModel(clientId int, coborrowersIdSlice []int) *IdsViewModel {
	idsvm := new(IdsViewModel)
	idsvm.SetClientId(clientId)
	idsvm.SetCoBorrowersIdSlice(coborrowersIdSlice)
	return idsvm
}

func (idsViewModel *IdsViewModel) GetClientId() int {
	return idsViewModel.ClientId
}

func (idsViewModel *IdsViewModel) GetCoborrowersIdSlice() []int {
	return idsViewModel.CoborrowersIdSlice
}

func (idsViewModel *IdsViewModel) validateClientId() bool {

	res := true

	clientId := idsViewModel.GetClientId()
	if clientId < 0 {

		idsViewModel.validation.AddMessage(Validation.GenMessage("clientId", "Is negative"))
		res = false

	}

	for _, id := range idsViewModel.CoborrowersIdSlice {

		if clientId == id {

			idsViewModel.validation.AddMessage(Validation.GenMessage("clientId", "matches the co-borrower id"))
			res = false

		}

	}

	return res
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (idsViewModel *IdsViewModel) validateCoBorrowerIdSliceUnique() bool {
	ids := idsViewModel.CoborrowersIdSlice
	ids_unique := unique(ids)

	if len(ids) != len(ids_unique) {
		return false
	}

	return true
}

func (idsViewModel *IdsViewModel) validateCoBorrowerIdSlice() bool {

	res := true

	for _, id := range idsViewModel.CoborrowersIdSlice {

		if id < 0 {

			idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", "Is negative"))
			res = false
		}

	}

	if !idsViewModel.validateCoBorrowerIdSliceUnique() {

		idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", "matches the co-borrower id"))
		res = false

	}

	return res
}

func (idsViewModel *IdsViewModel) Validate() bool {
	if idsViewModel.validateCoBorrowerIdSlice() && idsViewModel.validateClientId() {
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
	idsViewModel.validateCoBorrowerIdSlice()
	idsViewModel.validateClientId()
	return &idsViewModel.validation
}

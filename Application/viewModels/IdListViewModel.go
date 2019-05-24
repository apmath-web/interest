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
	clientId := idsViewModel.GetClientId()
	if clientId < 0 {
		idsViewModel.validation.AddMessage(Validation.GenMessage("clientId", "Is negative"))
		return false
	}

	for _, id := range idsViewModel.CoborrowersIdSlice {
		if clientId == id {
			idsViewModel.validation.AddMessage(Validation.GenMessage("clientId", "matches the co-borrower id"))
			return false
		}
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

	ids := idsViewModel.CoborrowersIdSlice
	for id1 := 0; id1 < len(ids); id1++ {
		for id2 := id1; id2 < len(ids); id2++ {
			if ids[id1] == ids[id2] && id1 != id2 {
				idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", "Is not unique"))
				return false
			}
		}

	}

	return true
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
	return &idsViewModel.validation
}

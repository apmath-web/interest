package viewModels

import (
	"testing"
)

func TestIdsViewModel_GetValidation(t *testing.T) {

	// check different coborrowers
	var vms = []IdsViewModel{
		IdsViewModel{
			ClientId:           1,
			CoborrowersIdSlice: []int{2, 3, 4},
		},

		IdsViewModel{
			ClientId:           1,
			CoborrowersIdSlice: []int{'0', '3'},
		},

		IdsViewModel{
			ClientId:           1,
			CoborrowersIdSlice: []int{},
		},
	}

	for _, value := range vms {
		if !value.Validate() {
			t.Error("Expected true, got", value.Validate())
		}
	}

	// check equal values client id and coborrower's id
	vms[0].SetClientId(2)

	if vms[0].Validate() {
		t.Error("Expected false, got", vms[0].Validate())
	}

	// check negative client id
	vms[0].SetClientId(-2)

	if vms[0].Validate() {
		t.Error("Expected false, got", vms[0].Validate())
	}

	// check equal values coborrowers id
	vms[0].SetCoBorrowersIdSlice([]int{2, 2})

	if vms[0].Validate() {
		t.Error("Expected false, got", vms[0].Validate())
	}

	// check negative coborrower's id
	vms[0].SetCoBorrowersIdSlice([]int{-2})

	if vms[0].Validate() {
		t.Error("Expected false, got", vms[0].Validate())
	}

}

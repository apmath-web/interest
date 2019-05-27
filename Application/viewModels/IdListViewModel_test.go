package viewModels

import (
	"testing"
)

func TestIdsViewModelPositiveCase_GetValidation(t *testing.T) {

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
		if !value.GetValidation().Empty() {
			t.Error("Expected true, got", value.GetValidation().Empty())
		}
	}

}

func TestIdsViewModelNegativeCase_GetValidation(t *testing.T) {

	var vms = []IdsViewModel{

		// check equal values client id and coborrower's id
		IdsViewModel{
			ClientId:           1,
			CoborrowersIdSlice: []int{1, 3, 4},
		},

		// check equal values coborrowers id
		IdsViewModel{
			ClientId:           1,
			CoborrowersIdSlice: []int{2, 2},
		},

		// check negative client id
		IdsViewModel{
			ClientId:           -1,
			CoborrowersIdSlice: []int{},
		},

		// check negative coborrower's id
		IdsViewModel{
			ClientId:           1,
			CoborrowersIdSlice: []int{-2},
		},

		// check negative coborrower's id and client id
		IdsViewModel{
			ClientId:           -1,
			CoborrowersIdSlice: []int{-2},
		},

		// check negative coborrower's id and client id
		IdsViewModel{
			ClientId:           -1,
			CoborrowersIdSlice: []int{-1},
		},
	}

	var errSlice = []string{

		"clientId: matches the co-borrower id",

		"coBorrowers: matches the co-borrower id",

		"clientId: Is negative",

		"coBorrowers: Is negative",

		"coBorrowers: Is negativeclientId: Is negative",

		"coBorrowers: Is negativeclientId: Is negativeclientId: matches the co-borrower id",
	}

	for i := 0; i < len(vms); i++ {

		var validatin = vms[i].GetValidation()

		if !validatin.Empty() {

			var allMessagesString string

			errMess := validatin.GetMessages()

			//t.Error(len(errMess))

			for _, message := range errMess {
				stringMess := message.GetStringMessage()
				allMessagesString += stringMess
			}

			if allMessagesString != errSlice[i] {

				t.Error("Expected " + errSlice[i] + " got" + allMessagesString)

			}

		}

	}
}

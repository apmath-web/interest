package models

import "github.com/apmath-web/interests/Domain"

type IdsDomainModel struct {
	ClientId           int
	CoborrowersIdSlice []int
}

func (i *IdsDomainModel) GetClientId() int {
	return i.ClientId
}

func (i *IdsDomainModel) GetCoborrowersIdSlice() []int {
	return i.CoborrowersIdSlice
}

func GenIds(clientId int, coborrowersIdSlice []int) Domain.IdsDomainModelInterface {
	idsModel := new(IdsDomainModel)
	idsModel.ClientId = clientId
	idsModel.CoborrowersIdSlice = coborrowersIdSlice
	return idsModel
}

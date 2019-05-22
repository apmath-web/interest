package models

import "github.com/apmath-web/interests/Domain"

type InterestsDomainModel struct {
	interest float64
}

func (ex *InterestsDomainModel) GetInterest() float64 {
	return ex.interest
}

func (ex *InterestsDomainModel) SetInterest(interest float64) {
	ex.interest = interest
}

func GenInterestsDomainModel(interest float64) Domain.InterestsInterface {
	dm := new(InterestsDomainModel)
	dm.interest = interest
	return dm
}
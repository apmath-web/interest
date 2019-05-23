package models

import (
	"github.com/apmath-web/interests/Domain"
)

type PersonDomainModel struct {
	FirstName string
	LastName  string
	SumWage   int
}

func (person *PersonDomainModel) GetFirstName() string {
	return person.FirstName
}

func (person *PersonDomainModel) GetLastName() string {
	return person.LastName
}

func (person *PersonDomainModel) GetSumWage() int {
	return person.SumWage
}



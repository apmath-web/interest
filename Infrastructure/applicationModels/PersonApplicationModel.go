package applicationModels

import (
	"encoding/json"
)

type JsonPerson struct {
	FirstName string                `json:"firstName"`
	LastName  string                `json:"lastName"`
	Jobs      []JobApplicationModel `json:"jobs"`
}

type PersonApplicationModel struct {
	JsonPerson
}

func (person *PersonApplicationModel) GetFirstName() string {
	return person.FirstName
}

func (person *PersonApplicationModel) GeLastName() string {
	return person.LastName
}

func (person *PersonApplicationModel) GetJobs() []JobApplicationModel {
	return person.Jobs
}

func (person *PersonApplicationModel) UnmarshalJSON(b []byte) error {
	tmpPerson := JsonPerson{}
	err := json.Unmarshal(b, &tmpPerson)
	if err != nil {
		return err
	}
	person.JsonPerson = tmpPerson
	return nil
}

func (person *PersonApplicationModel) GetSumWage() int {
	var sum int
	for _, value := range person.Jobs {
		sum += value.Wage
	}
	return sum
}

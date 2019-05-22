package Domain


type PersonDomainModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetSumWage() int
}

type IdsDomainModelInterface interface {
	GetClientId() int
	GetCoborrowersIdSlice() []int
}

type InterestsInterface interface {
	GetInterest() float64
	SetInterest(amount float64)
}

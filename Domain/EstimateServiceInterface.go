package Domain

type EstimateServiceInterface interface {
	Calculate(persons []PersonDomainModelInterface) (InterestsInterface, error)
}

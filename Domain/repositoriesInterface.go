package Domain

type repositoryInterface interface {
	GetModel(id int) HelloWorldApplicationModel
	PutModel(model HelloWorldApplicationModel) int
}

type ClientFetchInterface interface {
	Fetch(id int) (PersonDomainModelInterface, error)
}

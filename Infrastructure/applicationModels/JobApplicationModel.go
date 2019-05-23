package applicationModels

type JobApplicationModel struct {
	Name string
	Wage int
}

func (job *JobApplicationModel) GetName() string {
	return job.Name
}

func (job *JobApplicationModel) GetWage() int {
	return job.Wage
}



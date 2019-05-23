package applicationModels

type JobApplicationModel struct {
	Name string `json:"name"`
	Wage int `json:"wage"`
}

func (job *JobApplicationModel) GetName() string {
	return job.Name
}

func (job *JobApplicationModel) GetWage() int {
	return job.Wage
}

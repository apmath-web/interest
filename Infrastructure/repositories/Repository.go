package repositories

import (
	"github.com/apmath-web/interests/Domain"
)

type Repository struct {
	models         map[int]Domain.HelloWorldApplicationModel
	numberOfModels int
}

func GenRepository() *Repository {
	repo := &Repository{make(map[int]Domain.HelloWorldApplicationModel), 0}
	return repo
}

func (r *Repository) GetModel(id int) Domain.HelloWorldApplicationModel {
	helloWorld, ok := r.models[id]
	if ok {
		return helloWorld
	}
	return nil
}

func (r *Repository) PutModel(model Domain.HelloWorldApplicationModel) int {
	r.numberOfModels++
	r.models[r.numberOfModels] = model
	return r.numberOfModels
}

var Repo = GenRepository()

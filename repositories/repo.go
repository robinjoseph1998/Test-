package repositories

import (
	model "Test/models"
	"fmt"
)

type Repo struct {
	model model.DataModel
}

func NewRepo() *Repo {

	return &Repo{}

}

func (r *Repo) SaveName(name string) (string, error) {
	r.model.Name = name
	fmt.Println("MODEL", r.model.Name)
	return r.model.Name, nil
}

func (r *Repo) GetName() (string, error) {
	fmt.Println("Name", r.model.Name)
	return r.model.Name, nil
}

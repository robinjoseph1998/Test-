package usecase

import (
	"Test/repositories"
	"errors"
)

type UseCase struct {
	repo repositories.Repo
}

func NewUecase(r repositories.Repo) *UseCase {
	return &UseCase{repo: r}
}

func (rr UseCase) ExecuteAddname(name string) (string, error) {

	name, err := rr.repo.SaveName(name)
	if err != nil {
		return "", errors.New("can't add the name")
	}
	return name, nil
}

func (rr UseCase) ExecuteShowName() (string, error) {
	name, err := rr.repo.GetName()
	if err != nil {
		return "", errors.New("can't fetch the name")
	}
	return name, nil
}

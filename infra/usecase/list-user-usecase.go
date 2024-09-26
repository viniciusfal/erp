package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type ListUserUseCase struct {
	repository repository.UserRepository
}

func NewListUserUseCase(repo repository.UserRepository) ListUserUseCase {
	return ListUserUseCase{
		repository: repo,
	}
}

func (uu *ListUserUseCase) GetUsers() ([]model.User, error) {
	return uu.repository.GetUsers()
}

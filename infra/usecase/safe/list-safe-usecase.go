package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type ListSafeUseCase struct {
	repository repository.SafeRepository
}

func NewListSafesUseCase(repo repository.SafeRepository) ListSafeUseCase {
	return ListSafeUseCase{
		repository: repo,
	}
}

func (su *ListSafeUseCase) GetSafes() ([]model.Safe, error) {
	return su.repository.GetSafes()
}

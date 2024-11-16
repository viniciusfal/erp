package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type CreateSafeUseCase struct {
	repository repository.SafeRepository
}

func NewCreateSafeUseCase(repo repository.SafeRepository) CreateSafeUseCase {
	return CreateSafeUseCase{
		repository: repo,
	}
}

func (su *CreateSafeUseCase) CreateSafe(safe model.Safe) (model.Safe, error) {
	safe.Active = true

	_, err := su.repository.CreateSafe(safe)
	if err != nil {
		return model.Safe{}, err
	}

	return safe, nil
}

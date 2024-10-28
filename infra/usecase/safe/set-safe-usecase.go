package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SetSafeUseCase struct {
	repository repository.SafeRepository
}

func NewSetSafeUseCase(repo repository.SafeRepository) SetSafeUseCase {
	return SetSafeUseCase{
		repository: repo,
	}
}

func (su *SetSafeUseCase) SetSafe(safe *model.Safe) (*model.Safe, error) {
	safe, err := su.repository.SetSafe(safe)
	if err != nil {
		return nil, err
	}

	return safe, nil
}

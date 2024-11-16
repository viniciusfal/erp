package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SetActiveUseCase struct {
	repository repository.SafeRepository
}

func NewSetActiveUseCase(repo repository.SafeRepository) SetActiveUseCase {
	return SetActiveUseCase{
		repository: repo,
	}
}

func (su *SetActiveUseCase) SetActive(safe model.Safe) (*model.Safe, error) {
	_, err := su.repository.SetInativeSafe(safe)
	if err != nil {
		return nil, err
	}

	return &safe, nil
}

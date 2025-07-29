package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type UpdateConfigUseCase struct {
	repository repository.ConfigRepository
}

func NewUpdateConfigUseCase(repo repository.ConfigRepository) UpdateConfigUseCase {
	return UpdateConfigUseCase{
		repository: repo,
	}
}

func (uc *UpdateConfigUseCase) UpdateConfig(config *model.Config) (*model.Config, error) {
	updatedConfig, err := uc.repository.UpdateConfig(config)
	if err != nil {
		return nil, err
	}

	return updatedConfig, nil
} 
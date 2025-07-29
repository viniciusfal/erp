package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetConfigUseCase struct {
	repository repository.ConfigRepository
}

func NewGetConfigUseCase(repo repository.ConfigRepository) GetConfigUseCase {
	return GetConfigUseCase{
		repository: repo,
	}
}

func (gc *GetConfigUseCase) GetConfig() (*model.Config, error) {
	config, err := gc.repository.GetConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
} 
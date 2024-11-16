package usecase

import (
	"time"

	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetSafeByDateUseCase struct {
	repository repository.SafeRepository
}

func NewGetSafesByDateUseCase(repo repository.SafeRepository) GetSafeByDateUseCase {
	return GetSafeByDateUseCase{
		repository: repo,
	}
}

func (su *GetSafeByDateUseCase) GetSafesByDate(startDate time.Time, endDate time.Time) ([]*model.Safe, error) {
	safes, err := su.repository.GetSafeByDate(startDate, endDate)
	if err != nil {
		return nil, err
	}

	if startDate.After(endDate) {
		panic("A data final não pode ser menor que a inicial")
	}

	return safes, nil
}

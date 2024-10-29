package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetMetasByMonthUseCase struct {
	repository repository.MetaRepository
}

func NewMetasByMonthUseCase(repo repository.MetaRepository) GetMetasByMonthUseCase {
	return GetMetasByMonthUseCase{
		repository: repo,
	}
}

func (mu *GetMetasByMonthUseCase) GetMetaByMonth(month string) (*model.Meta, error) {
	meta, err := mu.repository.GetMetaByMonth(month)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetMetasUseCase struct {
	repository repository.MetaRepository
}

func NewGetMetasUseCase(repo repository.MetaRepository) GetMetasUseCase {
	return GetMetasUseCase{
		repository: repo,
	}
}

func (mu *GetMetasUseCase) GetMetas() ([]model.Meta, error) {
	metas, err := mu.repository.GetMetas()
	if err != nil {
		return nil, err
	}

	return metas, nil
}

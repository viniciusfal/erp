package usecase

import (
	"time"

	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type MetaUseCase struct {
	repository repository.MetaRepository
}

func NewMetaUseCase(repo repository.MetaRepository) MetaUseCase {
	return MetaUseCase{
		repository: repo,
	}
}

func (mu *MetaUseCase) CreateMeta(meta model.Meta) (model.Meta, error) {
	meta.Month = time.Now().Month().String()

	_, err := mu.repository.CreateMeta(meta)
	if err != nil {
		return model.Meta{}, err
	}

	return meta, nil
}

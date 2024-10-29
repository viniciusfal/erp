package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SetMetaUseCase struct {
	repository repository.MetaRepository
}

func NewSetMetaUseCase(repo repository.MetaRepository) SetMetaUseCase {
	return SetMetaUseCase{
		repository: repo,
	}
}

func (mu *SetMetaUseCase) SetMeta(meta model.Meta) (string, error) {
	metaUpdated, err := mu.repository.SetMeta(meta)
	if err != nil {
		return "", err
	}

	return metaUpdated, nil
}

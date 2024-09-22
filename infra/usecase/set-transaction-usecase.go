package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SetTransactionUseCase struct {
	repository repository.TransactionRepository
}

func NewSetTransactionUseCase(repo repository.TransactionRepository) SetTransactionUseCase {
	return SetTransactionUseCase{
		repository: repo,
	}
}

func (tu *SetTransactionUseCase) SetTransaction(t *model.Transaction) (*model.Transaction, error) {
	transaction, err := tu.repository.SetTransaction(t)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

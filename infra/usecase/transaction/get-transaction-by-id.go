package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetTransactionIdUseCase struct {
	repository repository.TransactionRepository
}

func NewGetTransactionByIdUseCase(repo repository.TransactionRepository) GetTransactionIdUseCase {
	return GetTransactionIdUseCase{
		repository: repo,
	}
}

func (tu *GetTransactionIdUseCase) GetTransactionById(id_transaction string) (*model.Transaction, error) {
	transaction, err := tu.repository.GetTransactionById(id_transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

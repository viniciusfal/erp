package usecase

import (
	"time"

	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type TransactionUseCase struct {
	repository repository.TransactionRepository
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return TransactionUseCase{
		repository: repo,
	}
}

func (tu *TransactionUseCase) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	_, err := tu.repository.CreateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction.Created_at = time.Now()

	return transaction, nil
}

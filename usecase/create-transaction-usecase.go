package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/viniciusfal/finances/model"
	"github.com/viniciusfal/finances/repository"
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

	transaction.ID = uuid.NewString()
	transaction.Created_at = time.Now()

	return transaction, nil
}

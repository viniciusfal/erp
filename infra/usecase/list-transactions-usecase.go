package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type ListTransactionUseCase struct {
	repository repository.TransactionRepository
}

func NewListTransactionUseCase(repo repository.TransactionRepository) ListTransactionUseCase {
	return ListTransactionUseCase{
		repository: repo,
	}
}

func (tu *ListTransactionUseCase) GetTransactions() ([]model.Transaction, error) {
	return tu.repository.GetProducts()
}

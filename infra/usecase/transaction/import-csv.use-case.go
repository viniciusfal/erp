package usecase

import (
	"time"

	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type ImportCSVUseCase struct {
	repository repository.TransactionRepository
}

func NewImportCSCVUseCase(repo repository.TransactionRepository) ImportCSVUseCase {
	return ImportCSVUseCase{
		repository: repo,
	}
}

func (tu *ImportCSVUseCase) ImportCSV(transactions []model.Transaction) ([]model.Transaction, error) {
	for i := range transactions {
		transactions[i].Created_at = time.Now()
	}

	// Chamar o repositório com os argumentos necessários
	err := tu.repository.ImportCSV(transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

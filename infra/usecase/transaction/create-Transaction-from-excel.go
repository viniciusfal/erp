package usecase

import (
	"mime/multipart"
	"time"

	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type TransactionImport struct {
	repository repository.TransactionRepository
}

func NewTransactionImport(repo repository.TransactionRepository) TransactionImport {
	return TransactionImport{
		repository: repo,
	}
}

func (tu *TransactionImport) CreateTransactionFromExcel(transaction model.Transaction, file multipart.File, fileHeader *multipart.FileHeader) (model.Transaction, error) {

	transaction.Created_at = time.Now()

	// Chamar o repositório com os argumentos necessários
	_, err := tu.repository.CreateTransaction(transaction, file, fileHeader)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

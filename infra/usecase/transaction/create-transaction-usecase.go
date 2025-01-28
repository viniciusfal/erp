package usecase

import (
	"mime/multipart"
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

func (tu *TransactionUseCase) CreateTransaction(transaction model.Transaction, file multipart.File, fileHeader *multipart.FileHeader) (model.Transaction, error) {

	transaction.Created_at = time.Now()

	// Ajustar o valor de Pay caso não seja agendado
	if !transaction.Scheduling {
		transaction.Pay = true
	}

	// Chamar o repositório com os argumentos necessários
	_, err := tu.repository.CreateTransaction(transaction, file, fileHeader)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

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
	// Define a data de criação
	transaction.Created_at = time.Now()

	// Ajustar o valor de Pay caso não seja agendado
	if !transaction.Scheduling {
		transaction.Pay = true
	}

	// Verifica se é uma transação de conta a pagar/receber
	if transaction.DueDate != nil && transaction.Status != nil {
		// Define o status inicial como "aberto" se não for fornecido
		if *transaction.Status == "" {
			status := "aberto"
			transaction.Status = &status
		}

		// Define o número da parcela e o total de parcelas, se necessário
		if transaction.Installment == nil {
			installment := 1
			transaction.Installment = &installment
		}
		if transaction.TotalInstallments == nil {
			totalInstallments := 1
			transaction.TotalInstallments = &totalInstallments
		}
	}

	// Chamar o repositório com os argumentos necessários
	_, err := tu.repository.CreateTransaction(transaction, file, fileHeader)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

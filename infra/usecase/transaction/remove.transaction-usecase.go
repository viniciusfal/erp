package usecase

import (
	"github.com/viniciusfal/erp/infra/repository"
)

type RemoveTransactionUseCase struct {
	repository repository.TransactionRepository
}

func NewRemoveTransactionUseCase(repo repository.TransactionRepository) RemoveTransactionUseCase {

	return RemoveTransactionUseCase{
		repository: repo,
	}
}

func (tu *RemoveTransactionUseCase) RemoveTransaction(transaction_id string) error {
	err := tu.repository.RemoveTransaction(transaction_id)
	if err != nil {
		return err
	}

	return nil

}

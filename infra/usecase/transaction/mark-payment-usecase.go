package usecase

import "github.com/viniciusfal/erp/infra/repository"

type MarkPaymentUseCase struct {
	repository repository.TransactionRepository
}

func NewMarkPaymentUseCase(repo repository.TransactionRepository) MarkPaymentUseCase {
	return MarkPaymentUseCase{
		repository: repo,
	}
}

func (tu *MarkPaymentUseCase) MarkPayment(transaction_id string) (string, error) {
	_, err := tu.repository.MarkPayment(transaction_id)
	if err != nil {
		return "nil", err
	}

	return transaction_id, nil
}

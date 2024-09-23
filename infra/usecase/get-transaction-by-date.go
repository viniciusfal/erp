package usecase

import (
	"time"

	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetTransactionByDateUseCase struct {
	repository repository.TransactionRepository
}

func NewGetTransactionByDateUseCase(repo repository.TransactionRepository) GetTransactionByDateUseCase {
	return GetTransactionByDateUseCase{
		repository: repo,
	}
}

func (tu *GetTransactionByDateUseCase) GetTransactionByDate(startDate time.Time, endDate time.Time) ([]*model.Transaction, error) {
	transactions, err := tu.repository.GetTransactionsByDate(startDate, endDate)
	if err != nil {
		return []*model.Transaction{}, err
	}

	if startDate.After(endDate) {
		panic("A data final não pode ser menor que a inicial")
	}

	return transactions, nil
}

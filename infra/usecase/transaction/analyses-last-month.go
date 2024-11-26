package usecase

import (
	"github.com/viniciusfal/erp/infra/repository"
)

type AnalysesTransactionUseCase struct {
	repository repository.TransactionRepository
}

func NewAnalysesTransactionUseCase(repo repository.TransactionRepository) AnalysesTransactionUseCase {
	return AnalysesTransactionUseCase{
		repository: repo,
	}
}

func (tu *AnalysesTransactionUseCase) GetTransactionByMonth() (int, int, int, error) {

	totalEntries, totalOutcomes, totalBalance, err := tu.repository.GetTransactionGrowthByMonth()
	if err != nil {
		return 0, 0, 0, err
	}

	return totalEntries, totalOutcomes, totalBalance, nil
}

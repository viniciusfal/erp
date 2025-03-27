package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetLowThirdyDays struct {
	transactionRepository repository.TransactionRepository
}

func NewLGetLowThirdyDays(repository repository.TransactionRepository) GetLowThirdyDays {
	return GetLowThirdyDays{
		transactionRepository: repository,
	}
}

func (tr *GetLowThirdyDays) GetTransactionByDate(ctx *gin.Context) {
	status := ctx.Query("status")

	transactions, err := tr.transactionRepository.GetLast30DaysTransactions(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

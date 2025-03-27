package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetLowLastSevenDays struct {
	transactionRepository repository.TransactionRepository
}

func NewLowGetLastSevenDays(repository repository.TransactionRepository) GetLowLastSevenDays {
	return GetLowLastSevenDays{
		transactionRepository: repository,
	}
}

func (tr *GetLowLastSevenDays) GetTransactionByDate(ctx *gin.Context) {
	status := ctx.Query("status")

	transactions, err := tr.transactionRepository.GetLast7DaysTransactions(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

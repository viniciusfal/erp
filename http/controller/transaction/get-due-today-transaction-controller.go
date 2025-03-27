package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetDueTodayController struct {
	transactionRepository repository.TransactionRepository
}

func NewGetDueTodayController(repository repository.TransactionRepository) GetDueTodayController {
	return GetDueTodayController{
		transactionRepository: repository,
	}
}

func (tr *GetDueTodayController) GetTransactionToday(ctx *gin.Context) {
	transactions, err := tr.transactionRepository.GetDueTodayTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetDueLastThirdyDaysController struct {
	transactionRepository repository.TransactionRepository
}

func NewLGetDueLastThirdyDaysController(repository repository.TransactionRepository) GetDueLastThirdyDaysController {
	return GetDueLastThirdyDaysController{
		transactionRepository: repository,
	}
}

func (tr *GetDueLastThirdyDaysController) GetDueTransactions(ctx *gin.Context) {
	typeTransaction := ctx.Query("type")

	transactions, err := tr.transactionRepository.GetLast30DaysTransactionsDueDate(typeTransaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

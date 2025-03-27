package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetDueLastSevenMonthController struct {
	transactionRepository repository.TransactionRepository
}

func NewGetDueLastSevenMonthController(repository repository.TransactionRepository) GetDueLastSevenMonthController {
	return GetDueLastSevenMonthController{
		transactionRepository: repository,
	}
}

func (tr *GetDueLastSevenMonthController) GetTransactionByDate(ctx *gin.Context) {
	transactions, err := tr.transactionRepository.GetLast7DaysTransactionsDueDate()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetLowCurrentMonthController struct {
	transactionRepository repository.TransactionRepository
}

func NewGetLowCurrentMonth(repository repository.TransactionRepository) GetLowCurrentMonthController {
	return GetLowCurrentMonthController{
		transactionRepository: repository,
	}
}

func (tr *GetLowCurrentMonthController) GetTransactionCurrent(ctx *gin.Context) {
	status := ctx.Query("status")

	transactions, err := tr.transactionRepository.GetCurreentMonthtransactions(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

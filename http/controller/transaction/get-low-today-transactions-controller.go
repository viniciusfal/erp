package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetLowTodayController struct {
	transactionRepository repository.TransactionRepository
}

func NewLGetLowTodayController(repository repository.TransactionRepository) GetLowTodayController {
	return GetLowTodayController{
		transactionRepository: repository,
	}
}

func (tr *GetLowTodayController) GetLowBydateToday(ctx *gin.Context) {
	status := ctx.Query("status")

	transactions, err := tr.transactionRepository.GetTodayTransactions(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

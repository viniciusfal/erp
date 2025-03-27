package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetDueCurrentMonthController struct {
	transactionRepository repository.TransactionRepository
}

func NewGetDueCurrentMonthController(repository repository.TransactionRepository) GetDueCurrentMonthController {
	return GetDueCurrentMonthController{
		transactionRepository: repository,
	}
}

func (tr *GetDueCurrentMonthController) GetDueCurrent(ctx *gin.Context) {

	transactions, err := tr.transactionRepository.GetCurreentMonthtransactionsDueDate()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

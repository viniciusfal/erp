package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/usecase"
)

type ListTransactionController struct {
	listTransactionUseCase usecase.ListTransactionUseCase
}

func NewListTransactionController(usecase usecase.ListTransactionUseCase) ListTransactionController {
	return ListTransactionController{
		listTransactionUseCase: usecase,
	}
}

func (tc *ListTransactionController) GetTransactions(ctx *gin.Context) {
	transactions, err := tc.listTransactionUseCase.GetTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

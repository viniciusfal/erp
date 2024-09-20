package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/usecase"
)

type transactionController struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionController(usecase usecase.TransactionUseCase) transactionController {
	return transactionController{
		transactionUseCase: usecase,
	}
}

func (tc *transactionController) CreateTransaction(ctx *gin.Context) {
	var transaction model.Transaction

	err := ctx.BindJSON(&transaction)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTransaction, err := tc.transactionUseCase.CreateTransaction(transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTransaction)

}

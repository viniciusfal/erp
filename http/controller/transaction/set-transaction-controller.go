package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

type SetTransactionController struct {
	setTransactionUseCase usecase.SetTransactionUseCase
}

func NewSetTransactionController(usecase usecase.SetTransactionUseCase) SetTransactionController {
	return SetTransactionController{
		setTransactionUseCase: usecase,
	}
}

func (tc *SetTransactionController) SetTransaction(ctx *gin.Context) {
	id := ctx.Param("transactionId")

	var transaction model.Transaction

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	transaction.ID = id

	updatedTransaction, err := tc.setTransactionUseCase.SetTransaction(&transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedTransaction)
}

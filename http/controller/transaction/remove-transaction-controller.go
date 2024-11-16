package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

type RemoveTransactionController struct {
	removeTransactionUseCase usecase.RemoveTransactionUseCase
}

func NewRemoveTransactionController(usecase usecase.RemoveTransactionUseCase) RemoveTransactionController {
	return RemoveTransactionController{
		removeTransactionUseCase: usecase,
	}
}

func (tc *RemoveTransactionController) RemoveTransaction(ctx *gin.Context) {
	id := ctx.Param("transactionId")

	err := tc.removeTransactionUseCase.RemoveTransaction(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

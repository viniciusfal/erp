package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/usecase"
)

type GetTransactionByIdController struct {
	getTransactionByIdusecase usecase.GetTransactionIdUseCase
}

func NewGetTransactionByIdController(usecase usecase.GetTransactionIdUseCase) GetTransactionByIdController {
	return GetTransactionByIdController{
		getTransactionByIdusecase: usecase,
	}
}

func (tc *GetTransactionByIdController) GetTransactionById(ctx *gin.Context) {
	id := ctx.Param("transactionId")

	transaction, err := tc.getTransactionByIdusecase.GetTransactionById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if transaction == nil {
		ctx.JSON(http.StatusNotFound, "Transação não encontrada")
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

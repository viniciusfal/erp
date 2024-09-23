package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/usecase"
)

type TransactionController struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionController(usecase usecase.TransactionUseCase) TransactionController {
	return TransactionController{
		transactionUseCase: usecase,
	}
}

func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
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

	formattedPaymentDate := ""
	if insertedTransaction.Payment_date != nil {
		formattedPaymentDate = insertedTransaction.PaymentDateFormat()
	}

	response := gin.H{
		"transaction_id": insertedTransaction.ID,
		"title":          insertedTransaction.Title,
		"value":          insertedTransaction.Value,
		"type":           insertedTransaction.Type,
		"category":       insertedTransaction.Category,
		"scheduling":     insertedTransaction.Scheduling,
		"annex":          insertedTransaction.Annex,
		"payment_date":   formattedPaymentDate,
		"created_at":     insertedTransaction.Created_at,
		"updated_at":     insertedTransaction.Updated_at,
	}

	ctx.JSON(http.StatusCreated, response)

}

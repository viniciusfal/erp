package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
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

	if transaction.Payment_date != nil {
		formattedPaymentDate := transaction.Payment_date.Format("02/01/2006")
		parsedDate, _ := time.Parse("02/01/2006", formattedPaymentDate)
		transaction.Payment_date = &parsedDate
	}

	insertedTransaction, err := tc.transactionUseCase.CreateTransaction(transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTransaction)
}

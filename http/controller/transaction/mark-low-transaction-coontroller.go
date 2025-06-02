package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type MarkLowTransactionController struct {
	TransactionRepository repository.TransactionRepository
}

func NewMarkLowTransactionController(repository repository.TransactionRepository) MarkLowTransactionController {
	return MarkLowTransactionController{
		TransactionRepository: repository,
	}
}

func (tc *MarkLowTransactionController) MarkLow(ctx *gin.Context) {
	var request struct {
		Status       string     `json:"status" binding:"required"`
		Payment_date *time.Time `json:"payment_date"`
		Method       string     `json:"method"`
		NF           *string    `json:"nf"`
		Account      *string    `json:"account"`
	}

	transactionId := ctx.Param("transactionId")
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTransactionToLow, err := tc.TransactionRepository.MarkLowTransaction(transactionId, request.Status, request.Payment_date, request.Method, request.NF, request.Account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o status: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateTransactionToLow)

}

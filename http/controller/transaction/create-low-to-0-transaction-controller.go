package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type CreateLowTo0TransactionController struct {
	TransactionRepository repository.TransactionRepository
}

func NewCreateLowTo0TransactionController(repository repository.TransactionRepository) CreateLowTo0TransactionController {
	return CreateLowTo0TransactionController{
		TransactionRepository: repository,
	}
}

func (tc *CreateLowTo0TransactionController) CreateLowTo0(ctx *gin.Context) {
	var request struct {
		TotalValue        float64   `json:"total_value" binding:"required"`
		TotalInstallments int       `json:"total_installments" binding:"required"`
		Title             string    `json:"title" binding:"required"`
		Details           string    `json:"details"`
		Type              string    `json:"type" binding:"required"`
		InitialDueDate    time.Time `json:"initial_due_date" binding:"required"`
		PaymentDate       time.Time `json:"payment_date" binding:"required"`
		Status            string    `json:"status" binding:"required"`
		Category          string    `json:"category" binding:"required"`
		NF                string    `json:"nf"`
		Pay               bool      `json:"pay"`
		Annex             string    `json:"annex"`
		Method            string    `json:"method" binding:"required"`
		Account           string    `json:"account"`
	}

	// Parse do JSON da requisição
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Chama o método do repositório para criar as transações parceladas
	transactionIDs, err := tc.TransactionRepository.CreateLowTransaction(
		request.TotalValue,
		request.TotalInstallments,
		request.Title,
		request.Details,
		request.Type,
		request.InitialDueDate,
		request.PaymentDate,
		request.Status,
		request.Category,
		request.NF,
		true,
		request.Annex,
		request.Method,
		request.Account,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar as transações parceladas: " + err.Error()})
		return
	}

	// Retorna os IDs das transações criadas
	ctx.JSON(http.StatusCreated, gin.H{"transaction_ids": transactionIDs})
}

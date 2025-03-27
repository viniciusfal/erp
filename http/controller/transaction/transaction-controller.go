package controller

import (
	"net/http"
	"strconv"
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

	transaction.Title = ctx.PostForm("title")
	transaction.Type = ctx.PostForm("type")
	transaction.Category = ctx.PostForm("category")
	schedulingStr := ctx.PostForm("scheduling")
	transaction.Details = ctx.PostForm("details")
	transaction.Method = ctx.PostForm("method")
	transaction.Nf = ctx.PostForm("nf")
	transaction.Account = ctx.PostForm("account")

	if supplierID := ctx.PostForm("supplier_id"); supplierID != "" {
		transaction.SupplierID = &supplierID
	} else {
		transaction.SupplierID = nil
	}

	valueStr := ctx.PostForm("value")
	if valueStr != "" {
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'value' deve ser um número"})
			return
		}
		transaction.Value = value
	}

	if annex := ctx.PostForm("annex"); annex != "" {
		transaction.Annex = &annex
	} else {
		transaction.Annex = nil
	}

	if schedulingStr != "" {
		scheduling, err := strconv.ParseBool(schedulingStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'scheduling' deve ser um booleano"})
			return
		}
		transaction.Scheduling = scheduling
	}

	// Formata as datas
	if ctx.PostForm("payment_date") != "" {
		parsedDate, err := time.Parse("02/01/2006", ctx.PostForm("payment_date"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use o formato ISO 8601 (ex: 2006-01-02T15:04:05Z)"})
			return
		}
		transaction.Payment_date = &parsedDate
	}

	if ctx.PostForm("due_date") != "" {
		parsedDate, err := time.Parse("02/01/2006", ctx.PostForm("due_date"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use o formato ISO 8601 (ex: 2006-01-02T15:04:05Z)"})
			return
		}
		transaction.DueDate = &parsedDate
	}

	// Processa campos opcionais
	if ctx.PostForm("status") != "" {
		status := ctx.PostForm("status")
		transaction.Status = &status
	}

	if ctx.PostForm("installment") != "" {
		installment, err := strconv.Atoi(ctx.PostForm("installment"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar o número da parcela"})
			return
		}
		transaction.Installment = &installment
	}

	if ctx.PostForm("total_installments") != "" {
		totalInstallments, err := strconv.Atoi(ctx.PostForm("total_installments"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar o total de parcelas"})
			return
		}
		transaction.TotalInstallments = &totalInstallments
	}

	// Processa o upload do arquivo (opcional)
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar o arquivo"})
		return
	}

	// Chama o usecase para criar a transação
	insertedTransaction, err := tc.transactionUseCase.CreateTransaction(transaction, file, fileHeader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar a transação"})
		return
	}

	// Retorna a transação criada
	ctx.JSON(http.StatusCreated, insertedTransaction)
}

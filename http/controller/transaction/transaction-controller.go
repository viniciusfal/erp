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

	// Usar ShouldBind para capturar os dados do formulário
	err := ctx.ShouldBind(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar os dados da transação"})
		return
	}

	// Processa o upload do arquivo
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar o arquivo"})
		return
	}

	// Se a data de pagamento não for nil, converte corretamente
	if transaction.Payment_date != nil {
		parsedDate, err := time.Parse("02/01/2006", transaction.Payment_date.Format("02/01/2006"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao formatar a data"})
			return
		}
		transaction.Payment_date = &parsedDate
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

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

type TransactionFromExcel struct {
	TransactionImportUseCase usecase.TransactionImport
}

func NewTransactionFromExcel(usecase usecase.TransactionImport) TransactionFromExcel {
	return TransactionFromExcel{
		TransactionImportUseCase: usecase,
	}
}

func (ti *TransactionFromExcel) ImportTransactions(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não enviado"})
		return
	}
	defer file.Close()

	var transaction model.Transaction

	_, err = ti.TransactionImportUseCase.CreateTransactionFromExcel(transaction, file, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar transações"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transações importadas com sucesso"})

}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
	"github.com/viniciusfal/erp/services"
)

type ImportCSVController struct {
	importCSVUseCase usecase.ImportCSVUseCase
}

func NewImportSCVController(usecase usecase.ImportCSVUseCase) ImportCSVController {
	return ImportCSVController{
		importCSVUseCase: usecase,
	}
}

func (ic *ImportCSVController) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não encontrado"})
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao abrir o arquivo"})
		return
	}
	defer src.Close()

	transactions, err := services.ParseCSV(src)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar CSV"})
		return
	}

	_, err = ic.importCSVUseCase.ImportCSV(transactions)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao salvar transações do arquivo .CSV"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "Criado com sucesso."})

}

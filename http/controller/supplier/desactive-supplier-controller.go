package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type DesactiveSupplierController struct {
	SupplierRepository repository.SupplierRepository
}

func NewDesactiveSupplierController(repository repository.SupplierRepository) DesactiveSupplierController {
	return DesactiveSupplierController{
		SupplierRepository: repository,
	}
}

func (ds *DesactiveSupplierController) Desactive(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ds.SupplierRepository.Desactive(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao desativar fornecedor"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Fornecedor desativado com sucesso"})
}

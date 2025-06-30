package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type FindOneSupplierController struct {
	SupplierRepository repository.SupplierRepository
}

func NewFindOneSupplierController(repository repository.SupplierRepository) FindOneSupplierController {
	return FindOneSupplierController{
		SupplierRepository: repository,
	}
}

func (fs *FindOneSupplierController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")

	supplier, err := fs.SupplierRepository.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar fornecedor"})
		return
	}

	if supplier == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Fornecedor não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, supplier)
}

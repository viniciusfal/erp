package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SupplierController struct {
	SupplierRepository repository.SupplierRepository
}

func NewSupplierController(repository repository.SupplierRepository) SupplierController {
	return SupplierController{
		SupplierRepository: repository,
	}
}

// CreateSupplier is a function to create a new supplier
func (sc *SupplierController) CreateSupplier(ctx *gin.Context) {
	var supplier model.Supplier

	supplier.ID = uuid.New().String()

	err := ctx.ShouldBind(&supplier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar os dados do fornecedor"})
		return
	}

	id, err := sc.SupplierRepository.Create(&supplier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o fornecedor"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SetSupplierController struct {
	SupplierRepository repository.SupplierRepository
}

func NewSetSupplierController(repository repository.SupplierRepository) SetSupplierController {
	return SetSupplierController{
		SupplierRepository: repository,
	}
}

func (ss *SetSupplierController) SetSupplier(ctx *gin.Context) {
	id := ctx.Param("id")

	var supplier model.Supplier

	err := ctx.ShouldBindJSON(&supplier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	supplier.ID = id

	err = ss.SupplierRepository.Update(&supplier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"supplier atualizado": supplier})
}

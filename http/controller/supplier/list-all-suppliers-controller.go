package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type ListAllSupplierController struct {
	SupplierRepository repository.SupplierRepository
}

func NewListAllSupplierController(repository repository.SupplierRepository) ListAllSupplierController {
	return ListAllSupplierController{
		SupplierRepository: repository,
	}
}

func (ls *ListAllSupplierController) ListSupplier(ctx *gin.Context) {
	suppliers, err := ls.SupplierRepository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, suppliers)

}

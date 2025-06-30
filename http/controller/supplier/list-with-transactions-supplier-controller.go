package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type ListWithTransactionController struct {
	SupplierRepository repository.SupplierRepository
}

func NewSupplierWithTransactionsController(repository repository.SupplierRepository) ListWithTransactionController {
	return ListWithTransactionController{
		SupplierRepository: repository,
	}
}

func (ls *ListWithTransactionController) ListSupplier(ctx *gin.Context) {
	suppliers, err := ls.SupplierRepository.FindAllWithTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, suppliers)

}

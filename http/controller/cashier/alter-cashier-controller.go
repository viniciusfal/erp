package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type AlterCashierController struct {
	cashierRepository repository.CashierRepository
}

func NewAlterCashierController(repository repository.CashierRepository) AlterCashierController {
	return AlterCashierController{
		cashierRepository: repository,
	}
}

func (cc *AlterCashierController) AlterCashierStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var cashier model.Cashier

	err := ctx.ShouldBindJSON(&cashier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	cashier.ID = id

	updateCashier, err := cc.cashierRepository.AlterCashier(cashier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updateCashier)

}

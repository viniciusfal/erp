package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SetCashierController struct {
	cashierRepository repository.CashierRepository
}

func NewSetCashierController(repository repository.CashierRepository) SetCashierController {
	return SetCashierController{
		cashierRepository: repository,
	}
}

func (cc *SetCashierController) SetCashier(ctx *gin.Context) {
	id := ctx.Param("id")
	var cashier model.Cashier

	err := ctx.ShouldBindJSON(&cashier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	cashier.ID = id

	updateCashier, err := cc.cashierRepository.SetCashier(cashier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updateCashier)

}

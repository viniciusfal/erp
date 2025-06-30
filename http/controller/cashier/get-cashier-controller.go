package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetCashierController struct {
	cashierRepository repository.CashierRepository
}

func NewGetCashierController(repository repository.CashierRepository) GetCashierController {
	return GetCashierController{
		cashierRepository: repository,
	}
}

func (cc *GetCashierController) GetAllCashier(ctx *gin.Context) {

	cashiers, err := cc.cashierRepository.GetCashier()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, cashiers)

}

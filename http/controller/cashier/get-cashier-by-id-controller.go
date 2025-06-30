package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetCashierByIDController struct {
	cashierRepository repository.CashierRepository
}

func NewGetCashierByIDController(repository repository.CashierRepository) GetCashierByIDController {
	return GetCashierByIDController{
		cashierRepository: repository,
	}
}

func (cc *GetCashierByIDController) GetCashierByID(ctx *gin.Context) {
	id := ctx.Param("id")

	cashier, err := cc.cashierRepository.GetCashierById(&id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, cashier)
}

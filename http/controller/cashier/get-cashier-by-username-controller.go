package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetCashierByUsername struct {
	cashierRepository repository.CashierRepository
}

func NewGetCashierByUsername(repository repository.CashierRepository) GetCashierByUsername {
	return GetCashierByUsername{
		cashierRepository: repository,
	}
}

func (cc *GetCashierByUsername) GetCashierByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	cashier, err := cc.cashierRepository.GetCashierByUserName(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, cashier)
}

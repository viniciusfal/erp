package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type CreateCashierController struct {
	createCashierRepository repository.CashierRepository
}

func NewCreateCashierController(repository repository.CashierRepository) CreateCashierController {
	return CreateCashierController{
		createCashierRepository: repository,
	}
}

func (cc *CreateCashierController) CreateCashier(ctx *gin.Context) {
	var cashier model.Cashier

	err := ctx.BindJSON(&cashier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	inseertedCashier, err := cc.createCashierRepository.CreateCashier(&cashier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, inseertedCashier)
}

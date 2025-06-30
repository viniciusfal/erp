package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type CreateAccController struct {
	accRepository repository.AccountabilityRepository
}

func NewCreateAccController(repository repository.AccountabilityRepository) CreateAccController {
	return CreateAccController{
		accRepository: repository,
	}
}

func (ac *CreateAccController) CreateAcc(ctx *gin.Context) {
	var acc model.Accountability

	err := ctx.BindJSON(&acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertAcc, err := ac.accRepository.CreateAcc(&acc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertAcc)
}

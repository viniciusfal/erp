package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SetAccController struct {
	accRepository repository.AccountabilityRepository
}

func NewSetAccController(repository repository.AccountabilityRepository) SetAccController {
	return SetAccController{
		accRepository: repository,
	}
}

func (ac *SetAccController) SetAcc(ctx *gin.Context) {
	id := ctx.Param("id")
	var accountability model.Accountability

	err := ctx.ShouldBindJSON(&accountability)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	accountability.ID = id

	updateAccountability, err := ac.accRepository.SetAcc(&accountability)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updateAccountability)

}

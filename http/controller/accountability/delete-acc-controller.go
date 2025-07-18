package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type DeleteAccController struct {
	accRepository repository.AccountabilityRepository
}

func NewDeleteAccController(repository repository.AccountabilityRepository) DeleteAccController {
	return DeleteAccController{
		accRepository: repository,
	}
}

func (ac *DeleteAccController) DeleteAcc(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id é obrigatório"})
		return
	}
	err := ac.accRepository.DeleteAcc(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
} 
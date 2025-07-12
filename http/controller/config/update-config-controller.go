package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/config"
)

type UpdateConfigController struct {
	updateConfigUseCase usecase.UpdateConfigUseCase
}

func NewUpdateConfigController(usecase usecase.UpdateConfigUseCase) UpdateConfigController {
	return UpdateConfigController{
		updateConfigUseCase: usecase,
	}
}

func (uc *UpdateConfigController) UpdateConfig(ctx *gin.Context) {
	var config model.Config

	err := ctx.ShouldBindJSON(&config)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	updatedConfig, err := uc.updateConfigUseCase.UpdateConfig(&config)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedConfig)
} 
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/config"
)

type GetConfigController struct {
	getConfigUseCase usecase.GetConfigUseCase
}

func NewGetConfigController(usecase usecase.GetConfigUseCase) GetConfigController {
	return GetConfigController{
		getConfigUseCase: usecase,
	}
}

func (gc *GetConfigController) GetConfig(ctx *gin.Context) {
	config, err := gc.getConfigUseCase.GetConfig()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, config)
} 
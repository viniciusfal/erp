package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/usecase"
)

type CreateMetaController struct {
	createMetauseCase usecase.MetaUseCase
}

func NewCeateMetaController(usecase usecase.MetaUseCase) CreateMetaController {
	return CreateMetaController{
		createMetauseCase: usecase,
	}
}

func (mc *CreateMetaController) CreateMeta(ctx *gin.Context) {
	var meta model.Meta

	err := ctx.BindJSON(&meta)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedMeta, err := mc.createMetauseCase.CreateMeta(meta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedMeta)

}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/meta"
)

type SetMetaController struct {
	setMetaUseCase usecase.SetMetaUseCase
}

func NewSetMetaController(usecase usecase.SetMetaUseCase) SetMetaController {
	return SetMetaController{
		setMetaUseCase: usecase,
	}
}

func (mc *SetMetaController) SetMeta(ctx *gin.Context) {
	id := ctx.Param("id")

	var meta model.Meta

	err := ctx.ShouldBindJSON(&meta)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	meta.ID = id

	updateMeta, err := mc.setMetaUseCase.SetMeta(meta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updateMeta)

}

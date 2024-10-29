package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/meta"
)

type GetMetasController struct {
	getMetasUseCase usecase.GetMetasUseCase
}

func NewGetMetasController(usecase usecase.GetMetasUseCase) GetMetasController {
	return GetMetasController{
		getMetasUseCase: usecase,
	}
}

func (mc *GetMetasController) GetMetas(ctx *gin.Context) {
	metas, err := mc.getMetasUseCase.GetMetas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, metas)
}

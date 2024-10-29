package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/meta"
)

type GetMetaByMonthController struct {
	getMEtaByMonthUseCase usecase.GetMetasByMonthUseCase
}

func NewGetMetaByMonthUseCase(usecase usecase.GetMetasByMonthUseCase) GetMetaByMonthController {
	return GetMetaByMonthController{
		getMEtaByMonthUseCase: usecase,
	}
}

func (mc *GetMetaByMonthController) GetMetaByMonth(ctx *gin.Context) {
	mounth := ctx.Param("month")

	meta, err := mc.getMEtaByMonthUseCase.GetMetaByMonth(mounth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, meta)
}

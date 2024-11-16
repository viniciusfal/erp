package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

type SetSafeController struct {
	setSafeUseCase usecase.SetSafeUseCase
}

func NewSetSafeController(usecase usecase.SetSafeUseCase) SetSafeController {
	return SetSafeController{
		setSafeUseCase: usecase,
	}
}

func (sc *SetSafeController) SetSafe(ctx *gin.Context) {
	id := ctx.Param("id")

	var safe model.Safe

	err := ctx.ShouldBindJSON(&safe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	safe.ID = id

	updatedSafe, err := sc.setSafeUseCase.SetSafe(&safe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedSafe)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

type SetActiveController struct {
	SetActiveUseCase usecase.SetActiveUseCase
}

func NewSetActiveController(usecase usecase.SetActiveUseCase) SetActiveController {
	return SetActiveController{
		SetActiveUseCase: usecase,
	}
}

func (sc *SetActiveController) SetActiveSafe(ctx *gin.Context) {
	id := ctx.Param("id")

	var safe model.Safe

	err := ctx.ShouldBindJSON(&safe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	safe.ID = id

	updateSafe, err := sc.SetActiveUseCase.SetActive(safe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updateSafe)
}

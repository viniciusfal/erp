package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

type ListSafeController struct {
	listSafeUseCase usecase.ListSafeUseCase
}

func NewListSafesController(usecase usecase.ListSafeUseCase) ListSafeController {
	return ListSafeController{
		listSafeUseCase: usecase,
	}
}

func (tc *ListSafeController) GetSafes(ctx *gin.Context) {
	safes, err := tc.listSafeUseCase.GetSafes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, safes)
}

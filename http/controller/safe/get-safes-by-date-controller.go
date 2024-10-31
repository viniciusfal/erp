package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

type GetSafesByDateController struct {
	GetSafesByDateUsecase usecase.GetSafeByDateUseCase
}

func NewGetSafesByDateController(usecase usecase.GetSafeByDateUseCase) GetSafesByDateController {
	return GetSafesByDateController{
		GetSafesByDateUsecase: usecase,
	}
}

func (sc *GetSafesByDateController) GetSafesByDate(ctx *gin.Context) {
	startDate := ctx.Param("startDate")
	endDate := ctx.Param("endDate")

	formattedStartDate, _ := time.Parse("2006-01-02", startDate)
	formattedEndDate, _ := time.Parse("2006-01-02", endDate)

	safes, err := sc.GetSafesByDateUsecase.GetSafesByDate(formattedStartDate, formattedEndDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, safes)
}

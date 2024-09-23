package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/usecase"
)

type GetTransactionByDateController struct {
	getTransactionByDateUseCase usecase.GetTransactionByDateUseCase
}

func NewGetTransactionByDateController(usecase usecase.GetTransactionByDateUseCase) GetTransactionByDateController {
	return GetTransactionByDateController{
		getTransactionByDateUseCase: usecase,
	}
}

func (tc *GetTransactionByDateController) GetTransactionByDate(ctx *gin.Context) {
	var request struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	err := ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	startFormattedDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		println(err)
		return
	}

	endFormattedDate, err := time.Parse("2006-01-02", request.EndDate)

	transactions, err := tc.getTransactionByDateUseCase.GetTransactionByDate(startFormattedDate, endFormattedDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)

}

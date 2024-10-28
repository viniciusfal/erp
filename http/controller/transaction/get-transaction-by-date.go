package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
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
	startDate := ctx.Param("startDate")
	endDate := ctx.Param("endDate")

	formattedStartDate, _ := time.Parse("02-01-2006", startDate)
	formattedEndDate, _ := time.Parse("02-01-2006", endDate)

	transactions, err := tc.getTransactionByDateUseCase.GetTransactionByDate(formattedStartDate, formattedEndDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

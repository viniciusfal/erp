package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

type AnalysesTransactionController struct {
	analysesTransactionUseCase usecase.AnalysesTransactionUseCase
}

func NewAnalysesTransactionController(usecase usecase.AnalysesTransactionUseCase) AnalysesTransactionController {
	return AnalysesTransactionController{
		analysesTransactionUseCase: usecase,
	}
}

func (tc *AnalysesTransactionController) GetTransactionByDate(ctx *gin.Context) {
	totalEntries, totalOutcomes, totalBalance, err := tc.analysesTransactionUseCase.GetTransactionByMonth()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_entries":  totalEntries,
		"total_outcomes": totalOutcomes,
		"total_balance":  totalBalance,
	})
}

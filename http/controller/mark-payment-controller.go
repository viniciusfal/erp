package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/usecase"
)

type MarkPaymentController struct {
	markPaymentUseCase usecase.MarkPaymentUseCase
}

func NewMarkPaymentUseController(usecase usecase.MarkPaymentUseCase) MarkPaymentController {
	return MarkPaymentController{
		markPaymentUseCase: usecase,
	}
}

func (tc *MarkPaymentController) MarkPayment(ctx *gin.Context) {
	transactionId := ctx.Param("transactionId")

	updatePayment, err := tc.markPaymentUseCase.MarkPayment(transactionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatePayment)

}

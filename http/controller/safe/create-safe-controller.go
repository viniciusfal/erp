package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

type CreateSafeController struct {
	safeUseCase usecase.CreateSafeUseCase
}

func NewCreateSafeController(usecase usecase.CreateSafeUseCase) CreateSafeController {
	return CreateSafeController{
		safeUseCase: usecase,
	}
}

func (sc *CreateSafeController) CreateSafe(ctx *gin.Context) {
	var safe model.Safe

	err := ctx.BindJSON(&safe)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	formattedPaymentDate := safe.Send_date.Format("02/01/2006")
	parsedDate, _ := time.Parse("02/01/2006", formattedPaymentDate)
	safe.Send_date = parsedDate

	insertedSafe, err := sc.safeUseCase.CreateSafe(safe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedSafe)
}

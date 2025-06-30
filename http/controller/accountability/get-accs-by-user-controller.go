package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetAccByUserController struct {
	accountabilityRepository repository.AccountabilityRepository
}

func NewGetAccByUserController(repository repository.AccountabilityRepository) GetAccByUserController {
	return GetAccByUserController{
		accountabilityRepository: repository,
	}
}

func (cc *GetAccByUserController) GetByUser(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")
	respId := ctx.Param("resp_id")

	formattedStartDate, _ := time.Parse("2006-01-02", startDate)
	formattedEndDate, _ := time.Parse("2006-01-02", endDate)

	acc, err := cc.accountabilityRepository.GetAccByUser(formattedStartDate, formattedEndDate, respId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

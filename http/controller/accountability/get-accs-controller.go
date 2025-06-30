package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetAccsController struct {
	accountabilityRepository repository.AccountabilityRepository
}

func NewGetAccsController(repository repository.AccountabilityRepository) GetAccsController {
	return GetAccsController{
		accountabilityRepository: repository,
	}
}

func (cc *GetAccsController) GetByDate(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")

	formattedStartDate, _ := time.Parse("2006-01-02", startDate)
	formattedEndDate, _ := time.Parse("2006-01-02", endDate)

	acc, err := cc.accountabilityRepository.GetAccBydate(formattedStartDate, formattedEndDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

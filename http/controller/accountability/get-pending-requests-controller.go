package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetPendingRequestsController struct {
	accountabilityRepository repository.AccountabilityRepository
}

func NewGetPendingRequestsController(repository repository.AccountabilityRepository) GetPendingRequestsController {
	return GetPendingRequestsController{
		accountabilityRepository: repository,
	}
}

func (cc *GetPendingRequestsController) GetPendingRequests(ctx *gin.Context) {
	status := ctx.Query("status")

	acc, err := cc.accountabilityRepository.GetPendingRequests(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

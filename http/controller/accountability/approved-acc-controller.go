package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type ApprovedAccController struct {
	accRepository repository.AccountabilityRepository
}

func NewApprovedAccController(repository repository.AccountabilityRepository) ApprovedAccController {
	return ApprovedAccController{
		accRepository: repository,
	}
}

func (ac *ApprovedAccController) ApprovedACC(ctx *gin.Context) {
	requestID := ctx.Param("requestId")
	adminID := ctx.Param("adminId")

	// Validação básica dos IDs
	if requestID == "" || adminID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDs não podem ser vazios"})
		return
	}

	result, err := ac.accRepository.ApproveAccountabilityChangeRequest(requestID, adminID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Falha ao aprovar solicitação",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

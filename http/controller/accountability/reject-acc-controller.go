package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type RejectAccController struct {
	accRepository repository.AccountabilityRepository
}

func NewRejectAccController(repository repository.AccountabilityRepository) RejectAccController {
	return RejectAccController{
		accRepository: repository,
	}
}

func (ac *RejectAccController) Reject(ctx *gin.Context) {

	requestID := ctx.Param("requestId")
	adminID := ctx.Param("adminId")

	// Validação básica dos IDs
	if requestID == "" || adminID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDs não podem ser vazios"})
		return
	}

	var requestBody struct {
		RejectionReason string `json:"rejection_reason" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Formato de requisição inválido",
			"details": err.Error(),
		})
		return
	}

	if requestBody.RejectionReason == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O motivo da rejeição é obrigatório"})
		return
	}

	result, err := ac.accRepository.RejectAccountabilityChangeRequest(requestID, adminID, requestBody.RejectionReason)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Falha ao rejeitar solicitação",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

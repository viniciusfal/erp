package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type ChangeACCRequest struct {
	accRepository repository.AccountabilityRepository
}

func NewChangeACCRequest(repository repository.AccountabilityRepository) *ChangeACCRequest {
	return &ChangeACCRequest{
		accRepository: repository,
	}
}

func (ac *ChangeACCRequest) ChangeACC(ctx *gin.Context) {
	var request model.AccountabilityChangeRequest

	// 1. Bind do JSON com tratamento de erro detalhado
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Erro no bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Formato de dados inválido",
			"details": err.Error(),
		})
		return
	}

	// 2. Validações básicas dos campos obrigatórios
	if request.OriginalAccountabilityID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "O ID da accountability original é obrigatório",
		})
		return
	}

	if request.RequestedBy == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "O identificador do solicitante é obrigatório",
		})
		return
	}

	// 3. Processamento no repositório
	createdRequest, err := ac.accRepository.ChangeRequest(&request)
	if err != nil {
		log.Printf("Erro no repositório: %v", err)

		// Verifica se é um erro de validação conhecido
		if isValidationError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Dados inválidos",
				"details": err.Error(),
			})
			return
		}

		// Erro genérico do servidor
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao processar solicitação",
			"details": "Por favor, tente novamente mais tarde",
		})
		return
	}

	// 4. Resposta de sucesso
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Solicitação de alteração criada com sucesso",
		"data":    createdRequest,
	})
}

// Função auxiliar para identificar erros de validação
func isValidationError(err error) bool {
	// Adicione aqui verificações para erros de validação específicos
	return strings.Contains(err.Error(), "violates foreign key constraint") ||
		strings.Contains(err.Error(), "violates not-null constraint")
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type UpdatePartnerController struct {
	partnerRepository repository.PartnerRepository
}

func NewUodatePartnerController(repository repository.PartnerRepository) UpdatePartnerController {
	return UpdatePartnerController{
		partnerRepository: repository,
	}
}

func (uc *UpdatePartnerController) Update(ctx *gin.Context) {
	partnerID := ctx.Param("id")
	if partnerID == "" {
		ctx.JSON(400, gin.H{"error": "Partner ID é obrigatorio"})
		return
	}

	var partner model.Partner
	if err := ctx.ShouldBindJSON(&partner); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	partner.ID = partnerID

	err := uc.partnerRepository.Update(&partner)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "falha ao atualizar partner"})
		return
	}

	ctx.JSON(200, partner)
}

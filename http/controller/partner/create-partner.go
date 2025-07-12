package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type CreatePartnerController struct {
	partnerRepository repository.PartnerRepository
}

func NewCreatePartnerController(repository repository.PartnerRepository) CreatePartnerController {
	return CreatePartnerController{
		partnerRepository: repository,
	}
}

func (cpc *CreatePartnerController) CreatePartner(ctx *gin.Context) {
	var partner model.Partner

	err := ctx.BindJSON(&partner)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	insertParcer, err := cpc.partnerRepository.Create(&partner)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "[Erro ao criar parceiro]"})
		return
	}

	ctx.JSON(201, insertParcer)
}

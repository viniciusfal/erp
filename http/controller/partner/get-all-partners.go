package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/repository"
)

type GetAllPartnersController struct {
	partnerRepository repository.PartnerRepository
}

func NewGetAllPartnersController(repository repository.PartnerRepository) GetAllPartnersController {
	return GetAllPartnersController{
		partnerRepository: repository,
	}
}

func (gc *GetAllPartnersController) GetAll(ctx *gin.Context) {
	partners, err := gc.partnerRepository.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "[Erro ao buscar parceiros]"})
		return
	}

	ctx.JSON(200, partners)
}

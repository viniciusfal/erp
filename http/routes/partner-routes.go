package routes

import (
	"github.com/gin-gonic/gin"
	factories "github.com/viniciusfal/erp/factories/partner"
)

func PartnerRoutes(router *gin.RouterGroup) {
	createPartnerController := factories.MakeCreatePartnerController()
	getallPartner := factories.MakeGetAllPartnersController()
	updatePartnerController := factories.MakeUpdatePartnerController()

	router.POST("/partner", createPartnerController.CreatePartner)
	router.GET("/partner", getallPartner.GetAll)
	router.PUT("/partner/:id", updatePartnerController.Update)
}

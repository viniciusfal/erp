package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/partner"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetAllPartnersController() controller.GetAllPartnersController {
	partnerRepository := repository.NewPartnerRepository(db.RunDB())
	getAllPartnersController := controller.NewGetAllPartnersController(partnerRepository)

	return getAllPartnersController
}

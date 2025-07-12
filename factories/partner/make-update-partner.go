package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/partner"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeUpdatePartnerController() controller.UpdatePartnerController {
	partnerRepository := repository.NewPartnerRepository(db.RunDB())
	updatePartnerController := controller.NewUodatePartnerController(partnerRepository)

	return updatePartnerController
}

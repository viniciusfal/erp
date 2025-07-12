package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/partner"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeCreatePartnerController() controller.CreatePartnerController {
	partnerRepository := repository.NewPartnerRepository(db.RunDB())
	createPartnerController := controller.NewCreatePartnerController(partnerRepository)

	return createPartnerController
}

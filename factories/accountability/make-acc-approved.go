package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/accountability"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeAccApproved() controller.ApprovedAccController {
	accRepository := repository.NewAccountabilityRepository(db.RunDB())
	accController := controller.NewApprovedAccController(accRepository)

	return accController
}

package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/accountability"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeAccReject() controller.RejectAccController {
	accRepository := repository.NewAccountabilityRepository(db.RunDB())
	accController := controller.NewRejectAccController(accRepository)

	return accController
}

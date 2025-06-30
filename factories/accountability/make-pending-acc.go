package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/accountability"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetPendingRequests() controller.GetPendingRequestsController {
	accRepository := repository.NewAccountabilityRepository(db.RunDB())
	accController := controller.NewGetPendingRequestsController(accRepository)

	return accController
}

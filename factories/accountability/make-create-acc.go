package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/accountability"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeCreateACC() controller.CreateAccController {
	accRepository := repository.NewAccountabilityRepository(db.RunDB())
	accController := controller.NewCreateAccController(accRepository)

	return accController
}

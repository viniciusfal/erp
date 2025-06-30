package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/accountability"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetAccsByDate() controller.GetAccsController {
	accRepository := repository.NewAccountabilityRepository(db.RunDB())
	accController := controller.NewGetAccsController(accRepository)

	return accController
}

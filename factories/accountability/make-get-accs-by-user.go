package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/accountability"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetAccsByUser() controller.GetAccByUserController {
	accRepository := repository.NewAccountabilityRepository(db.RunDB())
	accController := controller.NewGetAccByUserController(accRepository)

	return accController
}

package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/accountability"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeAccChangeRequest() controller.ChangeACCRequest {
	accRepository := repository.NewAccountabilityRepository(db.RunDB())
	accController := controller.NewChangeACCRequest(accRepository)

	return *accController
}

package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/cashier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeCreateCashier() controller.CreateCashierController {
	cashierRepository := repository.NewCashierRepository(db.RunDB())
	createCashierController := controller.NewCreateCashierController(cashierRepository)

	return createCashierController
}

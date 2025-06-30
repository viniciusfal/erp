package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/cashier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeSetCashier() controller.SetCashierController {
	cashierRepository := repository.NewCashierRepository(db.RunDB())
	setCashierController := controller.NewSetCashierController(cashierRepository)

	return setCashierController
}

package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/cashier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetAllCashier() controller.GetCashierController {
	cashierRepository := repository.NewCashierRepository(db.RunDB())
	getCashierController := controller.NewGetCashierController(cashierRepository)

	return getCashierController
}

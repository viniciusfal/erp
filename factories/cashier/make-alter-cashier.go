package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/cashier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeAlterCashierStatus() controller.AlterCashierController {
	cashierRepository := repository.NewCashierRepository(db.RunDB())
	alterCashierController := controller.NewAlterCashierController(cashierRepository)

	return alterCashierController
}

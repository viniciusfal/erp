package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/cashier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetCashierByUsername() controller.GetCashierByUsername {
	cashierRepository := repository.NewCashierRepository(db.RunDB())
	getCashierByUsername := controller.NewGetCashierByUsername(cashierRepository)

	return getCashierByUsername
}

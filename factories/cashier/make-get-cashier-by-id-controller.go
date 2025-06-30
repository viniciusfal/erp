package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/cashier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetcashierById() controller.GetCashierByIDController {
	cashierRepository := repository.NewCashierRepository(db.RunDB())
	getCashierByIDController := controller.NewGetCashierByIDController(cashierRepository)

	return getCashierByIDController
}

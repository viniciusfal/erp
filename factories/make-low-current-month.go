package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeGetLowCurrentMonth() controller.GetLowCurrentMonthController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MakeGetLowCurrentMonthController := controller.NewGetLowCurrentMonth(TransactionRepository)

	return MakeGetLowCurrentMonthController
}

package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeDueGetLastSevenDays() controller.GetDueLastSevenMonthController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MakeDueGetLastSevenDaysController := controller.NewGetDueLastSevenMonthController(TransactionRepository)

	return MakeDueGetLastSevenDaysController
}

package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeDueCurrentMonth() controller.GetDueCurrentMonthController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MakeDueCurrentMonthController := controller.NewGetDueCurrentMonthController(TransactionRepository)

	return MakeDueCurrentMonthController
}

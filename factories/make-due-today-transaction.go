package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeDueTodayTransaction() controller.GetDueTodayController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MakeGetDueTodayController := controller.NewGetDueTodayController(TransactionRepository)

	return MakeGetDueTodayController
}
